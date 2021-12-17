package router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"time"
)

var _ = Describe("Router unit tests", func() {

	Context("Header detection", func() {
		When("a request contains a given header", func() {
			It("should be available via the GetHeader() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Response {
					return request.Success(request.GetHeader("X-Custom-Header"))
				}(req)

				Expect(response.Content).To(Equal([]byte("exists")))
			})

			It("should be available via the GetHeaders() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Response {
					return request.Success(request.GetHeaders()["X-Custom-Header"][0])
				}(req)

				Expect(response.Content).To(Equal([]byte("exists")))
			})
		})
	})

	Context("URL parameter detection", func() {
		When("a URL parameter is detected", func() {
			It("should be detectable with the ArgExists function", func() {
				req := CreateRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})

				response := func(request Request) Response {
					if request.ArgExists("parameterOne") {
						return request.Success("found")
					}
					return request.Error("failed")
				}(req)

				Expect(response.Content).To(Equal([]byte("found")))
			})
		})

		When("a URL parameter is detected", func() {
			It("should be available via the GetArg method", func() {
				req := CreateRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})

				response := func(request Request) Response {
					if !request.ArgExists("parameterOne") {
						return request.Error("failed")
					}
					return request.Success(request.GetArg("parameterOne"))
				}(req)

				Expect(response.Content).To(Equal([]byte("exists")))
			})
		})
	})

	Context("Returns from multiple formats of variables", func() {
		When("the handler returns a string", func() {
			It("serves up a byte slice equivalent of the string", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success("string test")
				}(req)

				Expect(response.Content).To(Equal([]byte("string test")))
			})
		})

		When("the handler returns an integer", func() {
			It("serves up a byte slice equivalent of the integer", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(8)
				}(req)

				Expect(response.Content).To(Equal([]byte("8")))
			})
		})

		When("the handler returns a float", func() {
			It("serves up a byte slice equivalent of the float", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(5.6)
				}(req)

				Expect(response.Content).To(Equal([]byte("5.6")))
			})
		})

		When("the handler returns a struct", func() {
			It("serves up a byte slice equivalent of the struct, marshalled to json", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(struct{Status string}{Status: "success"})
				}(req)

				Expect(response.Content).To(Equal([]byte(`{"Status":"success"}`)))
			})
		})

		When("the handler returns a time", func() {
			It("serves up a byte slice equivalent of the time", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					t, _ := time.Parse("2006-01-02 15:04:05", "1981-12-03 13:00:00")
					return request.Success(t)
				}(req)

				Expect(response.Content).To(Equal([]byte("1981-12-03 13:00:00")))
			})
		})
	})
})
