package test_http_handler

import "net/http"

func handleGetFoo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("x-request-id", "cdd")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("FOO"))
}
