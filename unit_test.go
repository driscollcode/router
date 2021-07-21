package router

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRouting(t *testing.T) {
	router := Router{}
	router.Get("/", func(request Request) Response {
		return request.SuccessWithMsg("OK")
	})

	router.Get("/one/:two/:three", func(request Request) Response {
		return request.SuccessWithMsg(fmt.Sprintf("%s:%s:%s", "OK", request.GetArg("two"), request.GetArg("three")))
	})

	router.Post("/", func(request Request) Response {
		return request.SuccessWithMsg("OK:POST")
	})

	router.Get("/test/error", func(request Request) Response {
		return request.Error("OK:Fault")
	})

	router.Get("/route/with/[:optional]/[:parameters]", func(request Request) Response {
		return request.SuccessWithMsg(fmt.Sprintf("optional:%s-%s", request.GetArg("optional"), request.GetArg("parameters")))
	})

	r := httptest.NewRequest("GET", "http://example.org/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "OK" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/one/blah/works", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "OK:blah:works" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("POST", "http://example.org/", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "OK:POST" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/test/error", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 400 || w.Body.String() != "OK:Fault" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/this/does/not/exist", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 404 || w.Body.String() != "No provider could be found" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/route/with", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "optional:-" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/route/with/first", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "optional:first-" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}

	r = httptest.NewRequest("GET", "http://example.org/route/with/first/second", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, r)

	if w.Result().StatusCode != 200 || w.Body.String() != "optional:first-second" {
		t.Error("Unexpected result from request handler :", w.Result().StatusCode, ":", w.Body.String())
	}
}
