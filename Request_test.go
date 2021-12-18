package router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
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

		When("a URL parameter is not detected", func() {
			It("should cause the ArgExists method to return boolean false", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(request.ArgExists("unsetParameter"))
				}(req)

				Expect(response.Content).To(Equal([]byte("false")))
			})

			It("should be represented as the empty string when calling the GetArg method", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(request.GetArg("unsetParameter"))
				}(req)

				Expect(response.Content).To(Equal([]byte("")))
			})
		})
	})

	Context("Success responses", func() {
		When("the Success() method is called", func() {
			It("returns a HTTP 200 OK response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(nil)
				}(req)

				Expect(response.StatusCode).To(Equal(http.StatusOK))
			})
		})

		When("the Created() method is called", func() {
			It("returns a HTTP 201 Created response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Created(nil)
				}(req)

				Expect(response.StatusCode).To(Equal(http.StatusCreated))
			})
		})

		When("the Accepted() method is called", func() {
			It("returns a HTTP 202 Accepted response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Accepted(nil)
				}(req)

				Expect(response.StatusCode).To(Equal(http.StatusAccepted))
			})
		})

		When("the CustomResponse() method is called", func() {
			It("returns a custom HTTP status code response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.CustomResponse(999, nil)
				}(req)

				Expect(response.StatusCode).To(Equal(999))
			})
		})

		When("the SetHeader method is called", func() {
			It("causes a header to be added to the HTTP response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					request.SetHeader("Custom-Response-Header", "Set")
					return request.Success(nil)
				}(req)

				Expect(response.Headers["Custom-Response-Header"]).To(Equal("Set"))
			})
		})
	})

	Context("Returns from multiple formats of variables", func() {
		When("the handler returns nil", func() {
			It("serves up a byte slice containing nil", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(nil)
				}(req)

				Expect(response.Content).To(Equal([]byte(nil)))
			})
		})

		When("the handler returns a boolean true", func() {
			It("serves up a byte slice containing the string representation of a boolean true", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(true)
				}(req)

				Expect(response.Content).To(Equal([]byte("true")))
			})
		})

		When("the handler returns a boolean false", func() {
			It("serves up a byte slice containing the string representation of a boolean false", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Response {
					return request.Success(false)
				}(req)

				Expect(response.Content).To(Equal([]byte("false")))
			})
		})

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
