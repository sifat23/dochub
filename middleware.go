package main

import "net/http"

func methodHandler(method string, h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(errors) > 0 {
			// Combine all errors into a single string
			http.Error(w, "Errors: "+errors[0], http.StatusBadRequest)
			return
		}

		if r.Method != method {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		h(w, r)
	}
}
