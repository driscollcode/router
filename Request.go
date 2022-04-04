package router

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

type Request interface {
	ArgExists(name string) bool
	Body() []byte
	BodyError() error
	GetArg(name string) string
	GetHeader(header string) string
	GetHeaders() map[string][]string
	GetHost() string
	GetIP() string
	GetPostVariable(name string) string
	GetReferer() string
	GetURL() string
	GetUserAgent() string
	HasBody() bool
	HeaderExists(header string) bool
	PostVariableExists(name string) bool
}

func createRequest(method, path string, body []byte, params map[string]string) Request {
	return &request{input: httptest.NewRequest(method, path, bytes.NewReader(body)), args: params, URL: path}
}

func createRequestAdvanced(req *http.Request, params map[string]string) Request {
	return &request{input: req, args: params, Host: req.Host, URL: req.URL.Path, UserAgent: req.Header.Get("User-Agent")}
}

type request struct {
	input                *http.Request
	args                 map[string]string
	Host, URL, UserAgent string
	body                 struct {
		content   []byte
		error     error
		processed bool
	}
}

func (r *request) GetHost() string {
	return r.Host
}

func (r *request) GetUserAgent() string {
	return r.UserAgent
}

func (r *request) ArgExists(name string) bool {
	_, exists := r.args[name]
	return exists
}

func (r *request) GetArg(name string) string {
	arg, exists := r.args[name]
	if !exists {
		return ""
	}

	return arg
}

func (r *request) HeaderExists(header string) bool {
	return len(r.input.Header.Get(header)) > 0
}

func (r *request) GetHeader(header string) string {
	return r.input.Header.Get(header)
}

func (r *request) GetHeaders() map[string][]string {
	return r.input.Header
}

func (r *request) GetPostVariable(name string) string {
	err := r.input.ParseForm()
	if err != nil {
		return ""
	}

	return r.input.FormValue(name)
}

func (r *request) PostVariableExists(name string) bool {
	return len(r.GetPostVariable(name)) >= 1
}

func (r *request) GetURL() string {
	return r.input.URL.Path
}

func (r *request) GetIP() string {
	forwardedIPs := strings.Split(r.GetHeader("X-Forwarded-For"), ",")

	if len(forwardedIPs[0]) < 1 {
		return r.input.RemoteAddr
	}

	return forwardedIPs[0]
}

func (r *request) GetReferer() string {
	return r.input.Referer()
}

func (r *request) Body() []byte {
	r.processBody()
	return r.body.content
}

func (r *request) BodyError() error {
	r.processBody()
	return r.body.error
}

func (r *request) HasBody() bool {
	r.processBody()
	return len(r.body.content) > 0
}

func (r *request) processBody() {
	if r.body.processed {
		return
	}

	r.body.content, r.body.error = ioutil.ReadAll(r.input.Body)
	r.body.processed = true
}
