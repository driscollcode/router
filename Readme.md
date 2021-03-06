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

// User navigates to /user/bob in browser
func getUser(request router.Request) router.Response {
    if !request.ArgExists("name") {
        return request.Error("Name parameter is missing")
    }
    
    // Create a new user struct based on URL parameters
    user := struct{Name string}{Name: request.GetArg("name")}
    
    // Automatically send out a struct as the response body with a 200 status code
    return request.Success(user)
}
// HTTP response is 200:{"Name":"bob"}
```

### Specify An HTTP Status Code
```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/basic/example", basic)
	r.Serve(80)
}

func basic(request router.Request) router.Response {
    return request.Success(201, "this is a 201 response")
}
// HTTP response is 201:this is a 201 response
```

### Failed Response
```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/failure/example", fail)
	r.Serve(80)
}

func fail(request router.Request) router.Response {
    return request.Error("this is an example of a failure response")
}
// HTTP response is 400:this is an example of a failure response
```

### Failed Response With Custom HTTP Status Code
```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/failure/example", fail)
	r.Serve(80)
}

func fail(request router.Request) router.Response {
    return request.Error(500, "this is an example of a failure response")
}
// HTTP response is 500:this is an example of a failure response
```

### Request Functions

The following functions are defined on the ``Request`` struct and are available with each request.

* ``ArgExists(name string) bool`` - Does the named argument exists in the URL
* ``Body() []byte`` - Return the request body as a byte slice
* ``BodyError() error`` - Return an error if one occurred when fetching the request body
* ``GetArg(name string) string`` - Fetch the named argument from the URL
* ``GetBrowser() string`` - Guesstimate the browser from the request ``User-Agent`` header
* ``GetDeviceType() string`` - Guesstimate the device type from the request ``User-Agent`` header
* ``GetHeader(header string) string`` - Fetch the named request HTTP header
* ``GetHeaders() map[string][]string`` - Fetch a map of all request headers
* ``GetHost() string`` - Fetch the hostname of the request URL
* ``GetIP() string`` - Get the IP address of the request
* ``GetOperatingSystem() string`` - Guesstimate the operating system from the ``User-Agent`` header
* ``GetReferer() string`` - Get the HTTP Referer header from the request
* ``GetPostVariable(name string) string`` - Get the specified POST variable from the request
* ``GetUserAgent() string`` - Get the User Agent header from the request
* ``HeaderExists(header string) bool`` - Check if the specified header exists in the request
* ``HasBody() bool`` - Simple check to determine if the request has a body
* ``PostVariableExists(name string) bool`` - Check if the specified POST variable exists
* ``SetResponseHeader(key, value string)`` - Set a header for the request response

## Middleware

The router supports chains of handlers working together. The following example will output the line below.

`pre processed content - interesting content - post processed content`

As the request is handled, it is passed first to `postware` which runs the passed handler first, and then
adds it's own content. The passed handler is `preware` which adds it's content first, and then calls it's own
passed handler, in this case `myHandler`.

By changing the order in which handlers either modify a request themselves or pass it on to another handler,
you can control the order in which handlers add to the overal request.

The `request.Response()` method is used in `myHandler` - this will not set an HTTP status code if one is
already set - allowing you to defer this responsiblity to middleware. If no HTTP status code is set, `request.Response()`
will set the `HTTP 200 OK` code. You can easily override this anyway in any middleware by setting the first 
argument to your chosen status code.

In the example below, the `preware` is actually setting the status code to be `201`. If multiple handlers set a
status code, the last call to set a code is the one that goes to the browser.

```go
package main

import (
	"github.com/driscollcode/router"
)

func main() {
	myRouter := router.Router{}
	myRouter.Get("/", postware(preware(myHandler)))
	myRouter.Serve(80)
}

func preware(handler router.Handler) router.Handler {
	return func(request router.Request) router.Response {
		request.Success(201, "pre processed content - ")
		return handler(request)
	}
}

func postware(handler router.Handler) router.Handler {
	return func(request router.Request) router.Response {
		request = handler(request)
		return request.Response(" - post processed content")
	}
}

func myHandler(request router.Request) router.Response {
	return request.Response("interesting content")
}
```

### Response Functions

The following functions are part of the ``Request`` struct and can be the return value of a handler function.
The ``response`` parameters can be any inbuilt type (including ``[]byte``) or any struct that can be marshalled
to json. To respond with a specific HTTP status code, supply that as the first parameter to either the
``Error`` or ``Success`` function. The default status codes are shown below.

* ``Error(response ...interface{})`` - HTTP 400 (Bad Request) response with the supplied content
* ``Success(response ...interface{})`` - HTTP 200 OK response with the supplied content
* ``Response(response ...interface{})`` - Set response content without specifying an HTTP status code (see Middleware).

You can also perform a quick redirect with these functions.

* ``Redirect(destination string)`` - Perform a HTTP 302 redirect to the supplied destination
* ``PermanentRedirect(destination string)`` - Perform a HTTP 301 redirect to the supplied destination

## TLS And Self Signed Certificates

The router makes it easy to serve requests over TLS. Simply specify your key and certificate
as parameters to the `ServeWithTLS` method.

```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/failure/example", fail)
	r.ServeWithTLS(80, "--- my key ---", "--- my certificate ---")
}

func fail(request router.Request) router.Response {
    return request.Error(500, "this is an example of a failure response")
}
// HTTP response is 500:this is an example of a failure response
```

### Automatic Self Signed Certificates

If you leave the key and certificate parameters blank in your call to `ServeWithTLS`, the router will
automatically generate a self signed certificate for you

In the example below, a self singed certificate will be generated instantly and used to serve
requests.

```go
package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/failure/example", fail)
	r.ServeWithTLS(80, "", "")
}

func fail(request router.Request) router.Response {
    return request.Error(500, "this is an example of a failure response")
}
// HTTP response is 500:this is an example of a failure response
```

## Testing

### Mockgen

Have a look in the [testing-examples/mockgen](testing-examples/mockgen) folder for some working sample code which uses mock requests 
to unit test a handler.

### Mock Requests

Creating a request is very easy. Consider this example which passes a request to a handler and collects
the response.

```go
package main

import (
	"fmt"
	"github.com/driscollcode/router"
)

func main() {
	response := getUser(router.CreateRequest("GET", "/some/url", nil, map[string]string{"name": "John"}))
	fmt.Println(response.GetResponseStatusCode(), ":", string(response.GetResponseContent()))
}

func getUser(request router.Request) router.Response {
	if !request.ArgExists("name") {
		return request.Error("Name parameter is missing")
	}

	// fetch user from somewhere
	user := struct{ Name string }{Name: request.GetArg("name")}

	// Automatically send out a struct as the response body with a 200 status code
	return request.Success(user)
}
```

This example shows how easy it is to create a request and supply it to a handler.
