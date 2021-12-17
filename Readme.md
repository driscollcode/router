# Driscoll Router

A modern router which takes away as much boilerplate as possible, enabling rapid service development.

## Examples

### Basic Server

```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/user/:name", getUser)
	r.Serve(80)
}

func getUser(request router.Request) router.Response {
    if !request.ArgExists("name") {
        return request.Error("Name parameter is missing")
    }
    
    // fetch user from somewhere
    user := struct{Name string}{Name: "John"}
    
    // Automatically send out a struct as the response body with a 200 status code
    return request.Success(user)
}
```

### Request Response Functions

The following functions are part of the ``Request`` struct and can be the return value of a handler function.

* ``CustomResponse(statusCode int, response interface{})`` A custom response with the supplied HTTP status code and content
* ``Error(response interface{})`` - HTTP 400 (Bad Request) response with the supplied content
* ``Success(response interface{})`` - HTTP 200 OK response with the supplied content
* ``Created(response interface{})`` - HTTP 201 Created response with the supplied content
* ``Accepted(response interface{})`` - HTTP 202 Accepted response with the supplied content

You can also perform a quick redirect with these functions.

* ``Redirect(destination string)`` - Perform a HTTP 302 redirect to the supplied destination
* ``PermanentRedirect(destination string)`` - Perform a HTTP 301 redirect to the supplied destination

### Request Functions

The following functions are defined on the ``Request`` struct and are available with each request.

* ``ArgExists(name string) bool`` - Does the named argument exist in the URL
* ``Body() []byte`` - Return the request body as a byte slice
* ``BodyError() error`` - Return an error if one occurred when fetching the request body
* ``GetArg(name string) string`` - Fetch the named argument from the URL
* ``GetBrowser() string`` - Guesstimate the browser from the request ``User-Agent`` header
* ``GetDeviceType() string`` - Guesstimate the device type from the request ``User-Agent`` header
* ``GetHeader(header string) string`` - Fetch the named request HTTP header
* ``GetHeaders() map[string][]string`` - Fetch a map of all request headers
* ``GetIP() string`` - Get the IP address of the request
* ``GetOperatingSystem() string`` - Guesstimate the operating system from the ``User-Agent`` header
* ``GetReferer() string`` - Get the HTTP Referer header from the request
* ``GetPostVariable(name string) string`` - Get the specified POST variable from the request
* ``HasBody() bool`` - Simple check to determine if the request has a body
* ``PostVariableExists(name string) bool`` - Check if the specified POST variable exists
* ``SetHeader(key, value string)`` - Set a header for the request response