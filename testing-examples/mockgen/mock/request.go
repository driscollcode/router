package mock

import "github.com/driscollcode/router"

//go:generate mockgen -destination=mock-request.go -package=mock . Request
type Request interface {
	router.Request
}
