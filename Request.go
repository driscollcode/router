package router

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

// Use this function to simulate incoming HTTP / Cloud Function requests
func TestRequest(request *http.Request, params map[string]string) Request {
	return Request{input: request, output: nil, args: params, Host: request.Host, URL: request.URL.Path, UserAgent: request.Header.Get("User-Agent")}
}

type Request struct {
	input                *http.Request
	output               http.ResponseWriter
	args                 map[string]string
	Host, URL, UserAgent string
	body                 []byte
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

func (r *Request) GetHeader(header string) string {
	return r.input.Header.Get(header)
}

func (r *Request) GetHeaders() map[string][]string {
	return r.input.Header
}

func (r *Request) SetHeader(key, value string) {
	r.output.Header().Set(key, value)
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

func (r *Request) GetOperatingSystem() string {
	identifiers := make(map[string]string)

	identifiers["Android"] = "Android"
	identifiers["iPhone"] = "iOS"
	identifiers["iPad"] = "iOS"
	identifiers["Mac_PowerPC"] = "MacOS"
	identifiers["Macintosh"] = "MacOS"
	identifiers["Mac OS X"] = "MacOS"
	identifiers["Linux"] = "Linux"
	identifiers["Windows"] = "Windows"
	identifiers["FacebookLinkPreview"] = "facebookexternalhit"

	for key, os := range identifiers {
		if strings.Contains(r.GetHeader("User-Agent"), key) {
			return os
		}
	}

	return ""
}

func (r *Request) GetDeviceType() string {
	identifiers := make(map[string]string)

	identifiers["iPad"] = "iPad"
	identifiers["iPhone"] = "iPhone"
	identifiers["Tablet"] = "Tablet"
	identifiers["Android"] = "Android"
	identifiers["FacebookLinkPreview"] = "Bot"

	for key, device := range identifiers {
		if strings.Contains(r.GetHeader("User-Agent"), key) {
			if key == "Android" {
				if strings.Contains(r.GetHeader("User-Agent"), "Mobile") {
					return "Android Phone"
				} else {
					return "Android Tablet"
				}
			}

			return device
		}
	}

	return "Computer"
}

func (r *Request) GetBrowser() string {
	identifiers := make(map[string]string)

	identifiers["Chrome/"] = "Chrome"
	identifiers["Firefox"] = "Firefox"
	identifiers["Safari"] = "Safari"
	identifiers["OPR/"] = "Opera"
	identifiers["Opera/"] = "Opera"

	for key, browser := range identifiers {
		if strings.Contains(r.GetHeader("User-Agent"), key) {
			return browser
		}
	}

	return "Unknown"
}

func (r *Request) Referer() string {
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

func (r *Request) Error(msg string) Response {
	return Response{StatusCode: http.StatusBadRequest, Content: []byte(msg)}
}

func (r *Request) CustomResponse(statusCode int, msg string) Response {
	return Response{StatusCode: statusCode, Content: []byte(msg)}
}

func (r *Request) Success() Response {
	return Response{StatusCode: http.StatusOK}
}

func (r *Request) SuccessWithMsg(msg string) Response {
	return Response{StatusCode: http.StatusOK, Content: []byte(msg)}
}

func (r *Request) SuccessWithBytes(content []byte) Response {
	return Response{StatusCode: http.StatusOK, Content: content}
}

func (r *Request) SuccessWithJson(content interface{}) Response {
	bytes, err := json.Marshal(content)
	if err != nil {
		return Response{StatusCode: http.StatusInternalServerError, Content: []byte("Unable to convert server response to JSON")}
	}
	return Response{StatusCode: http.StatusOK, Content: bytes}
}

func (r *Request) Body() []byte {
	if len(r.body) > 0 {
		return r.body
	}

	bytes, err := ioutil.ReadAll(r.input.Body)

	if err != nil {
		return nil
	}

	r.body = bytes
	return bytes
}

func (r *Request) BodyError() error {
	body, err := ioutil.ReadAll(r.input.Body)

	if err != nil {
		return err
	}

	r.body = body
	return nil
}

func (r *Request) HasBody() bool {
	body, err := ioutil.ReadAll(r.input.Body)

	if err != nil || len(body) < 1 {
		return false
	}

	r.body = body
	return true
}
