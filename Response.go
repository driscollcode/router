package router

type Response struct {
	StatusCode int
	Headers    map[string]string
	Content    []byte
	Redirect   struct {
		DoRedirect  bool
		Destination string
	}
}
