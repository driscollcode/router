package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"time"
)

func CreateRequest(method, path string, body []byte, params map[string]string) Request {
	return Request{input: httptest.NewRequest(method, path, bytes.NewReader(body)), args: params, URL: path}
}

func CreateRequestAdvanced(request *http.Request, params map[string]string) Request {
	return Request{input: request, args: params, Host: request.Host, URL: request.URL.Path, UserAgent: request.Header.Get("User-Agent")}
}

type Request struct {
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

func (r *Request) ArgExists(name string) bool {
	_, exists := r.args[name]
	return exists
}

func (r *Request) GetArg(name string) string {
	arg, exists := r.args[name]
	if !exists {
		return ""
	}

	return arg
}

func (r *Request) HeaderExists(header string) bool {
	return len(r.input.Header.Get(header)) > 0
}

func (r *Request) GetHeader(header string) string {
	return r.input.Header.Get(header)
}

func (r *Request) GetHeaders() map[string][]string {
	return r.input.Header
}

func (r *Request) SetHeader(key, value string) {
	if len(r.responseHeaders) < 1 {
		r.responseHeaders = make(map[string]string)
	}
	r.responseHeaders[key] = value
}

func (r *Request) GetPostVariable(name string) string {
	err := r.input.ParseForm()
	if err != nil {
		return ""
	}

	return r.input.FormValue(name)
}

func (r *Request) PostVariableExists(name string) bool {
	return len(r.GetPostVariable(name)) >= 1
}

func (r *Request) GetURL() string {
	return r.input.URL.Path
}

func (r *Request) GetIP() string {
	forwardedIPs := strings.Split(r.GetHeader("X-Forwarded-For"), ",")

	if len(forwardedIPs[0]) < 1 {
		return r.input.RemoteAddr
	}

	return forwardedIPs[0]
}

func (r *Request) GetReferer() string {
	return r.input.Referer()
}

func (r *Request) Redirect(destination string) Response {
	response := Response{StatusCode: 302}
	response.Redirect.DoRedirect = true
	response.Redirect.Destination = destination
	return response
}

func (r *Request) PermanentRedirect(destination string) Response {
	response := Response{StatusCode: 301}
	response.Redirect.DoRedirect = true
	response.Redirect.Destination = destination
	return response
}

func (r *Request) Error(response interface{}) Response {
	return Response{StatusCode: http.StatusBadRequest, Headers: r.responseHeaders, Content: r.getResponseBody(response)}
}

func (r *Request) CustomResponse(statusCode int, response interface{}) Response {
	return Response{StatusCode: statusCode, Headers: r.responseHeaders, Content: r.getResponseBody(response)}
}

func (r *Request) Success(response interface{}) Response {
	return Response{StatusCode: http.StatusOK, Headers: r.responseHeaders, Content: r.getResponseBody(response)}
}

func (r *Request) Created(response interface{}) Response {
	return Response{StatusCode: http.StatusCreated, Headers: r.responseHeaders, Content: r.getResponseBody(response)}
}

func (r *Request) Accepted(response interface{}) Response {
	return Response{StatusCode: http.StatusAccepted, Headers: r.responseHeaders, Content: r.getResponseBody(response)}
}

func (r *Request) Body() []byte {
	r.processBody()
	return r.body.content
}

func (r *Request) BodyError() error {
	r.processBody()
	return r.body.error
}

func (r *Request) HasBody() bool {
	r.processBody()
	return len(r.body.content) > 0
}

func (r *Request) processBody() {
	if r.body.processed {
		return
	}

	r.body.content, r.body.error = ioutil.ReadAll(r.input.Body)
	r.body.processed = true
}

func (r *Request) getResponseBody(response interface{}) []byte {
	if response == nil {
		return nil
	}

	switch reflect.ValueOf(response).Kind() {
	case reflect.Struct:
		if _, ok := response.(time.Time); ok {
			return []byte(response.(time.Time).Format("2006-01-02 15:04:05"))
		}

		if bytes, err := json.Marshal(response); err == nil {
			return bytes
		}
	case reflect.Slice:
		if _, ok := response.([]byte); ok {
			return response.([]byte)
		}
	}

	return []byte(fmt.Sprintf("%v", response))
}
