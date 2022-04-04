package router

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unit Tests")
}

var _ = Describe("Router unit tests", func() {

	var router Router
	BeforeEach(func() {
		router = Router{}
	})

	Context("Basic request matching", func() {
		When("the router is matching a route to a handler", func() {
			When("the method is GET", func() {
				It("should pair a matching request to the appropriate GET handler", func() {
					router.Get("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("GET", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the GET handler if the method is not GET", func() {
					router.Get("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})

			When("the method is POST", func() {
				It("should pair a matching request to the appropriate POST handler", func() {
					router.Post("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the POST handler if the method is not POST", func() {
					router.Post("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("GET", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})

			When("the method is PUT", func() {
				It("should pair a matching request to the appropriate PUT handler", func() {
					router.Put("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("PUT", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the PUT handler if the method is not PUT", func() {
					router.Put("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})

			When("the method is PATCH", func() {
				It("should pair a matching request to the appropriate PATCH handler", func() {
					router.Patch("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("PATCH", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the PATCH handler if the method is not PATCH}", func() {
					router.Patch("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})

			When("the method is DELETE", func() {
				It("should pair a matching request to the appropriate DELETE handler", func() {
					router.Delete("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("DELETE", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the DELETE handler if the method is not DELETE", func() {
					router.Delete("/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})

			When("the method is arbritary", func() {
				It("should pair a matching request to the appropriate handler", func() {
					router.Route("CUSTOM", "/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("CUSTOM", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})

				It("should not pair a matching request to the handler if the method is not the right one", func() {
					router.Route("CUSTOM", "/", func(request Request) Response {
						return request.Success("OK")
					})

					r := httptest.NewRequest("POST", "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
				})
			})
		})

		When("the router is given a defined route", func() {
			When("the request matches the defined route", func() {
				It("should serve up the correct handler function", func() {
					router.Get("/defined/route", func(request Request) Response {
						return request.Success("matched route")
					})

					r := httptest.NewRequest("GET", "/defined/route", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("matched route"))
				})
			})

			When("the request does not match the defined route", func() {
				It("should serve up a 404 response", func() {
					router.Get("/defined/route", func(request Request) Response {
						return request.Success("matched route")
					})

					r := httptest.NewRequest("GET", "/undefined/route", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(404))
					Expect(w.Body.String()).To(Equal("No provider could be found"))
				})
			})
		})

		When("the router is matching a URL parameter", func() {
			It("should be able to find the parameter correctly", func() {
				router.Get("/url/param/:one", func(request Request) Response {
					if request.ArgExists("one") {
						return request.Success(request.GetArg("one"))
					}
					return request.Error("fault")
				})

				r := httptest.NewRequest("GET", "/url/param/working", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(200))
				Expect(w.Body.String()).To(Equal("working"))
			})
		})

		When("the router is given a URL prefix via the Root() method", func() {
			It("should store the URL prefix to attach to any other routes", func() {
				router.Get("/here", func(request Request) Response {
					return request.Success("OK")
				})

				router.Root("/my/url/prefix")

				r := httptest.NewRequest("GET", "/my/url/prefix/here", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("OK"))
			})
		})

		When("the router is serving up an error state", func() {
			It("should return an HTTP bad request code", func() {
				router.Get("/", func(request Request) Response {
					return request.Error(nil)
				})

				r := httptest.NewRequest("GET", "/", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusBadRequest))
			})
		})

		When("the router is given a NotFound function", func() {
			It("should serve that function whenever a URL is not defined as a route", func() {
				router.NotFound(func(request Request) Response {
					return request.Success("not found handled correctly")
				})

				r := httptest.NewRequest("GET", "/some/url/which/is/not/a/defined/route", nil)
				w := httptest.NewRecorder()
				router.ServeHTTP(w, r)

				Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal("not found handled correctly"))
			})
		})
	})
})
