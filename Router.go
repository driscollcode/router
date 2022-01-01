package router

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

type Router struct {
	routes   []route
	notFound handler
	root     string
}

func (r *Router) Get(path string, handler handler) {
	r.url("GET", path, handler)
}

func (r *Router) Post(path string, handler handler) {
	r.url("POST", path, handler)
}

func (r *Router) Put(path string, handler handler) {
	r.url("PUT", path, handler)
}

func (r *Router) Patch(path string, handler handler) {
	r.url("PATCH", path, handler)
}

func (r *Router) Delete(path string, handler handler) {
	r.url("DELETE", path, handler)
}

func (r *Router) Route(method, path string, handler handler) {
	r.url(method, path, handler)
}

func (r *Router) NotFound(handler handler) {
	r.notFound = handler
}

func (r *Router) Root(urlRoot string) {
	r.root = urlRoot
}

func (r *Router) url(method, path string, handler handler) {
	if len(r.routes) < 1 {
		r.routes = make([]route, 0)
	}

	r.routes = append(r.routes, route{Method: method, Path: path, Handler: handler})
}

func (rt *Router) Serve(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), rt)
}

func (rt *Router) ServeIP(ip string, port int) error {
	return http.ListenAndServe(fmt.Sprintf("%s:%d", ip, port), rt)
}

func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		rt.corsInjector(w)
		w.WriteHeader(http.StatusOK)
		return
	}

	foundHandler, params, err := rt.findHandler(r)
	if err != nil {

		if rt.notFound != nil {
			foundHandler = rt.notFound
		} else {
			w.WriteHeader(404)
			w.Write([]byte("No provider could be found"))
			return
		}
	}

	req := request{input: r, args: params, Host: r.Host, URL: r.URL.Path, UserAgent: r.Header.Get("User-Agent")}
	response := foundHandler(&req)

	if len(os.Getenv("BuildDate")) > 0 {
		w.Header().Set("X-Build-Date", os.Getenv("BuildDate"))
	}

	rt.corsInjector(w)

	if response.Redirect.DoRedirect {
		http.Redirect(w, r, response.Redirect.Destination, response.StatusCode)
		return
	}

	for key, val := range response.Headers {
		w.Header().Set(key, val)
	}

	w.WriteHeader(response.StatusCode)
	w.Write(response.Content)
}

func (rt *Router) findHandler(r *http.Request) (handler, map[string]string, error) {
	for _, route := range rt.routes {
		if !strings.EqualFold(r.Method, route.Method) {
			continue
		}

		match, args := rt.isAMatch(route.Path, r.URL.Path)

		if match {
			return route.Handler, args, nil
		}
	}

	return nil, nil, errors.New("no_handler")
}

func (r *Router) isAMatch(path, url string) (bool, map[string]string) {
	if len(r.root) > 0 {
		url = url[len(r.root):]
	}

	urlBits := strings.Split(strings.Trim(url, "/"), "/")
	pathBits := strings.Split(strings.Trim(path, "/"), "/")

	mandatoryBits := 0
	for _, bit := range pathBits {
		if strings.Contains(bit, "[") && strings.Contains(bit, "]") {
			mandatoryBits++
		}
	}

	if len(urlBits) < mandatoryBits || len(urlBits) > len(pathBits) {
		return false, nil
	}

	match := true
	args := make(map[string]string)
	for pos, bit := range pathBits {

		if len(urlBits) < pos+1 {
			if strings.Contains(bit, "[:") {
				break

			} else {
				args = nil
				match = false
				break
			}
		}

		if match && len(bit) > 1 && bit[0:1] == ":" {
			args[bit[1:]] = urlBits[pos]

		} else if match && len(bit) > 1 && bit[0:2] == "[:" {
			args[bit[2:len(bit)-1]] = urlBits[pos]

		} else {

			if urlBits[pos] != pathBits[pos] {
				args = nil
				match = false
				break
			}
		}
	}

	return match, args
}

func (r *Router) corsInjector(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
}
