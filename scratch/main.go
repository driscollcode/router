package main

import "github.com/driscollcode/router"

func main() {
	r := router.Router{}
	r.Get("/user/:name", getUser)
	r.Serve(80)
}

func getUser(call router.Call) router.Call {
	if !call.ArgExists("name") {
		return call.Error("Name parameter is missing")
	}

	// fetch user from somewhere
	user := struct{ Name string }{Name: call.GetArg("name")}

	// Automatically send out a struct as the response body with a 200 status code
	return call.Success(user)
}
