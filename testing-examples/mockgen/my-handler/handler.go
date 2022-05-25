package myHandler

import (
	"fmt"
	"github.com/driscollcode/router"
)

func Handler(request router.Request) router.Response {
	if !request.ArgExists("userId") {
		return request.Error("No user Id supplied")
	}
	return request.Success(fmt.Sprintf("%s processed", request.GetArg("userId")))
}
