package webmiddlewares

import (
	"fmt"
	"net/http"
)

type MiddleWare func(http.HandlerFunc) http.HandlerFunc

func ChainMiddleWares(fn http.HandlerFunc, mw ...MiddleWare) http.HandlerFunc {
	for _, m := range mw {
		fn = m(fn)
	}
	return fn
}

func LoggingMiddleWare() MiddleWare {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Logger Middleware", "r.URL", r.RequestURI)
			fn(w, r)
		}
	}
}

func MethodMiddleWare() MiddleWare {
	return func(fn http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Method Middleware", "r.Method", r.Method)
			fn(w, r)
		}
	}
}
