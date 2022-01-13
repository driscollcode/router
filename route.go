package router

type route struct {
	Method, Path string
	Handler      Handler
}