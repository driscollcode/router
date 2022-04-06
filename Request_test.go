package router

import (
	"bytes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("Router unit tests", func() {

	Context("Header detection", func() {
		When("a request contains a given header", func() {
			It("should be detectable via the HeaderExists() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")

				req := createRequestAdvanced(r, nil)
				Expect(req.HeaderExists("X-Custom-Header")).To(BeTrue())
			})

			It("should be available via the GetHeader() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := createRequestAdvanced(r, nil)

				Expect(req.GetHeader("x-custom-header")).To(Equal("exists"))
			})

			It("should be available via the GetHeaders() method", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("x-custom-header", "exists")
				req := createRequestAdvanced(r, nil)

				Expect(req.GetHeaders()["X-Custom-Header"]).To(Equal([]string{"exists"}))
			})
		})
	})

	Context("URL parameter detection", func() {
		When("a URL parameter is detected", func() {
			It("should be detectable with the ArgExists function", func() {
				req := createRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})
				Expect(req.ArgExists("parameterOne")).To(BeTrue())
			})

			It("should be available via the GetArg method", func() {
				req := createRequest("GET", "/", nil, map[string]string{"parameterOne": "exists"})
				Expect(req.GetArg("parameterOne")).To(Equal("exists"))
			})
		})

		When("a URL parameter is not detected", func() {
			It("should cause the ArgExists method to return boolean false", func() {
				req := createRequest("GET", "/", nil, nil)
				Expect(req.ArgExists("unsetParameter")).To(BeFalse())
			})

			It("should cause the ArgExists method to return boolean false even if it is populated blank", func() {
				req := createRequest("GET", "/", nil, map[string]string{"parameterOne": ""})
				Expect(req.ArgExists("parameterOne")).To(BeFalse())
			})

			It("should be represented as the empty string when calling the GetArg method", func() {
				req := createRequest("GET", "/", nil, nil)
				Expect(req.GetArg("unsetParameter")).To(Equal(""))
			})
		})
	})

	Context("Request information", func() {
		When("the GetURL method is called", func() {
			It("should return the url of the request", func() {
				req := createRequest("GET", "/this/is/the/url", nil, nil)
				Expect(req.GetURL()).To(Equal("/this/is/the/url"))
			})
		})

		When("the GetIP method is called", func() {
			It("should return the X-Forwarded-For header if it is present", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("X-Forwarded-For", "127.0.0.1")
				req := createRequestAdvanced(r, nil)
				Expect(req.GetIP()).To(Equal("127.0.0.1"))
			})

			It("should return the RemoteAddr request property if the X-Forwarded-For header is not present", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.RemoteAddr = "127.0.0.2"
				req := createRequestAdvanced(r, nil)
				Expect(req.GetIP()).To(Equal("127.0.0.2"))
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
				req := createRequestAdvanced(r, nil)

				Expect(req.PostVariableExists("post-data")).To(BeTrue())
			})

			It("should return false if the specified post variable is not present", func() {
				values := url.Values{}
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.PostVariableExists("post-data")).To(BeFalse())
			})

			It("should return false when HasBody() is called and there is no post body", func() {
				values := url.Values{}
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.HasBody()).To(BeFalse())
			})

			It("should return true when HasBody() is called and there is a post body", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.HasBody()).To(BeTrue())
			})

			It("should return a byte slice when the post body is requested via the Body() method", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.Body()).To(Equal([]byte("post-data=set")))
			})

			It("should return an empty byte slice via the Body() method when there is no post body", func() {
				values := url.Values{}
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.Body()).To(Equal([]byte("")))
			})

			It("should return a nil error when BodyError() is called as it is very hard to trip this error", func() {
				r := httptest.NewRequest("GET", "/", nil)
				req := createRequestAdvanced(r, nil)
				Expect(req.BodyError()).To(BeNil())
			})
		})

		When("the GetPostVariable method is called", func() {
			It("should return the requested post variable as a string if it is present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.GetPostVariable("post-data")).To(Equal("set"))
			})

			It("should return an empty string if the value is not present", func() {
				values := url.Values{}
				values.Set("post-data", "set")
				r := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
				r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
				req := createRequestAdvanced(r, nil)

				Expect(req.GetPostVariable("unset-post-variable")).To(Equal(""))
			})
		})
	})

	Context("Obtaining the Referer header", func() {
		When("the handler calls the GetReferer() method", func() {
			It("responds with the referer header when it is set", func() {
				r := httptest.NewRequest("GET", "/", nil)
				r.Header.Set("Referer", "https://example.org")
				req := createRequestAdvanced(r, nil)

				Expect(req.GetReferer()).To(Equal("https://example.org"))
			})
		})
	})
})
