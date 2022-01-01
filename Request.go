package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Request interface {
	ArgExists(name string) bool
	Body() []byte
	BodyError() error
	Error(response ...interface{}) Response
	GetArg(name string) string
	GetHeader(header string) string
	GetHeaders() map[string][]string
	GetIP() string
	GetPostVariable(name string) string
	GetReferer() string
	GetURL() string
	HasBody() bool
	HeaderExists(header string) bool
	PermanentRedirect(destination string) Response
	PostVariableExists(name string) bool
	Redirect(destination string) Response
	SetHeader(key, value string)
	Success(response ...interface{}) Response
}


func CreateRequest(method, path string, body []byte, params map[string]string) Request {
	return &request{input: httptest.NewRequest(method, path, bytes.NewReader(body)), args: params, URL: path}
}

func CreateRequestAdvanced(req *http.Request, params map[string]string) Request {
	return &request{input: req, args: params, Host: req.Host, URL: req.URL.Path, UserAgent: req.Header.Get("User-Agent")}
}

type request struct {
	input                *http.Request
	args                 map[string]string
	Host, URL, UserAgent string
	body struct {
		content []byte
		error error
		processed bool
	}
	responseHeaders map[string]string
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

func (r *request) SetHeader(key, value string) {
	if len(r.responseHeaders) < 1 {
		r.responseHeaders = make(map[string]string)
	}
	r.responseHeaders[key] = value
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

func (r *request) Redirect(destination string) Response {
	response := Response{StatusCode: 302}
	response.Redirect.DoRedirect = true
	response.Redirect.Destination = destination
	return response
}

func (r *request) PermanentRedirect(destination string) Response {
	response := Response{StatusCode: 301}
	response.Redirect.DoRedirect = true
	response.Redirect.Destination = destination
	return response
}

func (r *request) Error(response ...interface{}) Response {
	return Response{StatusCode: r.getStatusCode(http.StatusBadRequest, response...), Headers: r.responseHeaders, Content: r.getResponseBody(response...)}
}

func (r *request) Success(response ...interface{}) Response {
	return Response{StatusCode: r.getStatusCode(http.StatusOK, response...), Headers: r.responseHeaders, Content: r.getResponseBody(response...)}
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

func (r *request) getResponseBody(response ...interface{}) []byte {
	if response == nil {
		return nil
	}

	output := make([]byte, 0)
	for pos, piece := range response {
		if pos == 0 {
			if asInt, ok := piece.(int); ok && asInt >= 100 && asInt <= 999 {
				continue
			}
		}

		if len(output) > 0 {
			output = append(output, []byte("")...)
		}
		output = append(output, r.getContentAsByte(piece)...)
	}
	return output
}

func (r *request) getContentAsByte(content interface{}) []byte {
	switch reflect.ValueOf(content).Kind() {
	case reflect.Struct:
		if _, ok := content.(time.Time); ok {
			return []byte(fmt.Sprintf("%s", content.(time.Time).Format("2006-01-02 15:04:05")))
		}

		if bytes, err := json.Marshal(content); err == nil {
			return bytes
		}
	case reflect.Bool:
		return []byte(fmt.Sprint(content))
	case reflect.Slice:
		if bytes, ok := content.([]byte); ok {
			return bytes
		}
		return []byte(fmt.Sprint(content))
	case reflect.String:
		return []byte(content.(string))
	case reflect.Int:
		return []byte(strconv.Itoa(content.(int)))
	case reflect.Float64:
		return []byte(strconv.FormatFloat(content.(float64), 'f', -1, 64))
	}
	return nil
}

func (r *request) getStatusCode(defaultCode int, parts ...interface{}) int {
	if len(parts) < 1 {
		return defaultCode
	}

	userCode, err := strconv.Atoi(fmt.Sprintf("%v", parts[0]))
	if err != nil {
		return defaultCode
	}

	if userCode >= 100 && userCode <= 999 {
		return userCode
	}
	return defaultCode
}
