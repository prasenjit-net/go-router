package main

import (
	"go-router/eureka"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.Handle("/", &httputil.ReverseProxy{Director: eureka.Director})
	http.ListenAndServe(":8080", nil)
}
