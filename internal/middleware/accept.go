package middleware

import "net/http"

func Accept(next http.Handler) http.Handler {
	fn := func(resp http.ResponseWriter, req *http.Request) {
		if req.Header.Get("Accept") != "application/json" {
			http.Error(resp, "Not acceptable content-type", http.StatusNotAcceptable)
			return
		}

		next.ServeHTTP(resp, req)
	}

	return http.HandlerFunc(fn)
}
