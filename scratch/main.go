package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	getUser(getUser(router.CreateRequest("GET", "/", nil, nil)))
	r.Get("/user/:name", getUser)
	r.Serve(80)
}

func getUser(request router.Request) router.Response {
	if !request.ArgExists("name") {
		return request.Error("Name parameter is missing")
	}

	request.GetResponseWriter().Write([]byte("Pre Message\n"))

	// fetch user from somewhere
	user := struct{ Name string }{Name: request.GetArg("name")}

	// Automatically send out a struct as the response body with a 200 status code
	return request.Success(user)
}
