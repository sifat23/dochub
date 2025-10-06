package main

import (
	"dochub/bin"
	"net/http"
	"strings"
)

func methodHandler(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(bin.Errors) > 0 {
			http.Error(w, strings.Join(bin.Errors, ", "), http.StatusBadRequest)
			return
		}

		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}
