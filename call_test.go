package router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"time"
)

var _ = Describe("Call struct unit tests", func() {
	Context("Preparing a response", func() {
		When("setting a response header", func() {
			It("should appear in the call response struct", func() {
				call := call{}
				call.SetHeader("Custom-Header", "set")
				Expect(call.response.Headers["Custom-Header"]).To(Equal("set"))
			})
		})

		When("setting up a redirect", func() {
			It("should put the redirect into the response and set the appropriate properties", func() {
				call := call{}
				call.Redirect("http://example.org")

				Expect(call.response.StatusCode).To(Equal(http.StatusFound))
				Expect(call.response.Redirect.DoRedirect).To(BeTrue())
				Expect(call.response.Redirect.Destination).To(Equal("http://example.org"))
			})
		})

		When("setting up a permanent redirect", func() {
			It("should put the redirect into the response and set the appropriate properties", func() {
				call := call{}
				call.PermanentRedirect("http://example.org")

				Expect(call.response.StatusCode).To(Equal(http.StatusMovedPermanently))
				Expect(call.response.Redirect.DoRedirect).To(BeTrue())
				Expect(call.response.Redirect.Destination).To(Equal("http://example.org"))
			})
		})

		When("giving a success response", func() {
			When("supplying no arguments at all", func() {
				It("should set up a 200 OK response and no content", func() {
					call := call{}
					call.Success()

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("")))
				})
			})

			When("supplying a simple response", func() {
				It("should store the response in the content", func() {
					call := call{}
					call.Success("success-response")

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("success-response")))
				})
			})

			When("supplying a response that starts with an HTTP code", func() {
				It("should catch the status code and store that in the status code field and the response in the response field", func() {
					call := call{}
					call.Success(404, "couldn't see that")

					Expect(call.response.StatusCode).To(Equal(http.StatusNotFound))
					Expect(call.response.Content).To(Equal([]byte("couldn't see that")))
				})
			})
		})

		When("giving an error response", func() {
			When("no parameters are given at all", func() {
				It("should set the status code to 400 bad request and give no response body", func() {
					call := call{}
					call.Error()

					Expect(call.response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(call.response.Content).To(Equal([]byte("")))
				})
			})

			When("a string parameter is given", func() {
				It("will set the status code to 400 bad request and set the response body to the string parameter", func() {
					call := call{}
					call.Error("this is an error")

					Expect(call.response.StatusCode).To(Equal(http.StatusBadRequest))
					Expect(call.response.Content).To(Equal([]byte("this is an error")))
				})
			})

			When("an integer is given as the first parameter and a string is given as the second", func() {
				It("will take the integer as the http status code and the string as the response", func() {
					call := call{}
					call.Error(200, "this is an error")

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("this is an error")))
				})
			})
		})

		When("giving a generic response", func() {
			When("supplying no status code", func() {
				It("should return an HTTP 200 OK status code", func() {
					call := call{}
					call.Response("response-text")

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("response-text")))
				})
			})
		})

		When("working with the response system generally", func() {
			When("given an invalid http status code", func() {
				It("should fall back to the default for that method eg. HTTP 200 OK for success", func() {
					call := call{}
					call.Success(50000, "this is an invalid status code")

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("50000this is an invalid status code")))
				})
			})

			When("given a struct as a response", func() {
				It("should put the marshalled equivalent into the response body", func() {
					call := call{}
					call.Success(struct{ Name string }{Name: "Bob"})

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte(`{"Name":"Bob"}`)))
				})
			})

			When("given a time as a response", func() {
				It("should put the ISO version of the time into the response body", func() {
					call := call{}
					tm, _ := time.Parse("2006-01-02 15:04:05", "2022-01-01 12:00:00")
					call.Success(tm)

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("2022-01-01 12:00:00")))
				})
			})

			When("given a boolean as a response", func() {
				It("should give the textual representation of the boolean as part of the response body", func() {
					call := call{}
					call.Success(true)

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("true")))
				})
			})

			When("given a slice that converts easily to a slice of bytes (eg a string)", func() {
				It("should convert the slice into bytes", func() {
					call := call{}
					call.Success([]string{"One", "two"})

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("[One two]")))
				})
			})

			When("given a float", func() {
				It("converts it to a string", func() {
					call := call{}
					call.Success(1.2)

					Expect(call.response.StatusCode).To(Equal(http.StatusOK))
					Expect(call.response.Content).To(Equal([]byte("1.2")))
				})
			})
		})
	})
})
