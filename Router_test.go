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
		When("the router is matching a basic route with a specific method", func() {

			handlerFunc := func(request Request) Response {
				return request.Success("OK")
			}

			for _, method := range []string{"GET", "POST", "PUT", "PATCH", "DELETE"} {
				It("should match " + method + " requests", func() {
					switch method {
					case "GET":
						router.Get("/", handlerFunc)
					case "POST":
						router.Post("/", handlerFunc)
					case "PUT":
						router.Put("/", handlerFunc)
					case "PATCH":
						router.Patch("/", handlerFunc)
					case "DELETE":
						router.Delete("/", handlerFunc)
					}

					r := httptest.NewRequest(method, "/", nil)
					w := httptest.NewRecorder()
					router.ServeHTTP(w, r)

					Expect(w.Result().StatusCode).To(Equal(200))
					Expect(w.Body.String()).To(Equal("OK"))
				})
			}
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