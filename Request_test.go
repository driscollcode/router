package router

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"
)

var _ = Describe("Router unit tests", func() {

	Context("Header detection", func() {
		When("a request contains a given header", func() {
			It("should be detectable via the HeaderExists() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					switch request.HeaderExists("X-Custom-Header") {
					case true:
						return request.Success()
					default:
						return request.Error()
					}
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusOK))
			})

			It("should be available via the GetHeader() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetHeader("X-Custom-Header"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("exists")))
			})

			It("should be available via the GetHeaders() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetHeaders()["X-Custom-Header"][0])
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("exists")))
			})
		})
	})

	Context("URL parameter detection", func() {
		When("a URL parameter is detected", func() {
			It("should be detectable with the ArgExists function", func() {
				req := CreateRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})

				response := func(request Request) Request {
					if request.ArgExists("parameterOne") {
						return request.Success("found")
					}
					return request.Error("failed")
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("found")))
			})

			It("should be available via the GetArg method", func() {
				req := CreateRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})

				response := func(request Request) Request {
					if !request.ArgExists("parameterOne") {
						return request.Error("failed")
					}
					return request.Success(request.GetArg("parameterOne"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("exists")))
			})
		})

		When("a URL parameter is not detected", func() {
			It("should cause the ArgExists method to return boolean false", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(request.ArgExists("unsetParameter"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("false")))
			})

			It("should be represented as the empty string when calling the GetArg method", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(request.GetArg("unsetParameter"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})
		})
	})

	Context("Request information", func() {
		When("the GetURL method is called", func() {
			It("should return the url of the request", func() {
				req := CreateRequest("GET", "/this/is/the/url", nil, nil)

				response := func(request Request) Request {
					return request.Success(request.GetURL())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("/this/is/the/url")))
			})
		})

		When("the GetIP method is called", func() {
			It("should return the X-Forwarded-For header if it is present", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("X-Forwarded-For", "127.0.0.1")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetIP())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("127.0.0.1")))
			})

			It("should return the RemoteAddr request property if the X-Forwarded-For header is not present", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.RemoteAddr = "127.0.0.2"
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetIP())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("127.0.0.2")))
			})
		})
	})

	Context("POST data", func() {
		When("the PostVariableExists method is called", func() {
			It("should return true if the specified post variable is present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.PostVariableExists("post-data"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("true")))
			})

			It("should return false if the specified post variable is not present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.PostVariableExists("unset-post-variable"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("false")))
			})

			It("should return false when HasBody() is called and there is no post body", func() {
				values := url.Values{}
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.HasBody())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("false")))
			})

			It("should return true when HasBody() is called and there is a post body", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.HasBody())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("true")))
			})

			It("should return a byte slice when the post body is requested via the Body() method", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.Body())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("post-data=set")))
			})

			It("should return an empty byte slice via the Body() method when there is no post body", func() {
				values := url.Values{}
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.Body())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})

			It("should return a nil error when BodyError() is called as it is very hard to trip this error", func() {
				r := httptest.NewRequest("GET", "/", nil)
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.BodyError())
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})
		})

		When("the GetPostVariable method is called", func() {
			It("should return the requested post variable as a string if it is present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetPostVariable("post-data"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("set")))
			})

			It("should return an empty string if the value is not present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetPostVariable("unset-post-variable"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})
		})
	})

	Context("Success responses", func() {
		When("the Success() method is called", func() {
			It("returns a HTTP 200 OK response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success()
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusOK))
			})
		})

		When("the Success() method is called with a custom response code", func() {
			It("returns the user specified response code", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(http.StatusAccepted, "OK")
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusAccepted))
				Expect(response.GetResponse().Content).To(Equal([]byte("OK")))
			})
		})

		When("the Error() method is called", func() {
			It("returns a HTTP Bad Request status code", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Error()
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusBadRequest))
				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})
		})

		When("the Error() method is called with a custom response code", func() {
			It("returns the user specified response code", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Error(http.StatusUnauthorized, "OK")
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusUnauthorized))
				Expect(response.GetResponse().Content).To(Equal([]byte("OK")))
			})
		})

		When("the SetHeader method is called", func() {
			It("causes a header to be added to the HTTP response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					request.SetResponseHeader("Custom-Response-Header", "Set")
					return request.Success()
				}(req)

				Expect(response.GetResponse().Headers["Custom-Response-Header"]).To(Equal("Set"))
			})
		})
	})

	Context("Responses which perform a redirect", func() {
		When("the handler triggers a temporary redirect", func() {
			It("serves up a HTTP 302 status response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Redirect("/moved")
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(302))
				Expect(response.GetResponse().Redirect.Destination).To(Equal("/moved"))
			})
		})

		When("the handler triggers a permanent redirect", func() {
			It("serves up a HTTP 301 status response", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.PermanentRedirect("/moved permanently")
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(301))
				Expect(response.GetResponse().Redirect.Destination).To(Equal("/moved permanently"))
			})
		})
	})

	Context("Obtaining the Referer header", func() {
		When("the handler calls the GetReferer() method", func() {
			It("responds with the referer header when it is set", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("Referer", "https://example.org")
				req := CreateRequestAdvanced(r, nil)

				response := func(request Request) Request {
					return request.Success(request.GetReferer())
				}(req)

				Expect(response.GetResponse().StatusCode).To(Equal(http.StatusOK))
				Expect(response.GetResponse().Content).To(Equal([]byte("https://example.org")))
			})
		})
	})

	Context("Returns from multiple formats of variables", func() {
		When("the handler returns nil", func() {
			It("serves up a byte slice containing nil", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success()
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("")))
			})
		})

		When("the handler returns a boolean true", func() {
			It("serves up a byte slice containing the string representation of a boolean true", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(true)
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("true")))
			})
		})

		When("the handler returns a boolean false", func() {
			It("serves up a byte slice containing the string representation of a boolean false", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(false)
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("false")))
			})
		})

		When("the handler returns a string", func() {
			It("serves up a byte slice equivalent of the string", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success("string test")
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("string test")))
			})
		})

		When("the handler returns an integer", func() {
			It("serves up a byte slice equivalent of the integer", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(8)
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("8")))
			})
		})

		When("the handler returns a float", func() {
			It("serves up a byte slice equivalent of the float", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(5.6)
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("5.6")))
			})
		})

		When("the handler returns a struct", func() {
			It("serves up a byte slice equivalent of the struct, marshalled to json", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success(struct{Status string}{Status: "success"})
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte(`{"Status":"success"}`)))
			})
		})

		When("the handler returns a byte slice", func() {
			It("serves up the byte slice directly", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					return request.Success([]byte("byte slice content"))
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("byte slice content")))
			})
		})

		When("the handler returns a time", func() {
			It("serves up a byte slice equivalent of the time", func() {
				req := CreateRequest("GET", "/", nil, nil)

				response := func(request Request) Request {
					t, _ := time.Parse("2006-01-02 15:04:05", "1981-12-03 13:00:00")
					return request.Success(t)
				}(req)

				Expect(response.GetResponse().Content).To(Equal([]byte("1981-12-03 13:00:00")))
			})
		})
	})
})
