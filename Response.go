package router

type Response interface {
	Request
}

type response struct {
	statusCode int
	headers    map[string]string
	content    []byte
	redirect   struct {
		doRedirect  bool
		destination string
	}
}

func (r response) GetResponseStatusCode() int {
	return r.statusCode
}

func (r response) GetResponseHeaders() map[string]string {
	return r.headers
}

func (r response) GetResponseContent() []byte {
	return r.content
}

func (r response) GetResponseRedirect() string {
	if !r.redirect.doRedirect || len(r.redirect.destination) < 1 {
		return ""
	}
	return r.redirect.destination
}
