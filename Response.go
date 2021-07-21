package router

type Response struct {
	StatusCode int
	Content    []byte
	Redirect   struct {
		DoRedirect  bool
		Destination string
	}
}
