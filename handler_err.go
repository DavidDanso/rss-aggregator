package main

import (
	"net/http"
)

// function signature that you need to use of you want to define a HTTP handler
// from interface http.Handler
func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Something went grong")
}