package main

import (
	"net/http"
)

// function signature that you need to use of you want to define a HTTP handler
// from interface http.Handler
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}