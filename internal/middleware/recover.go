package middleware

import (
	"log"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	fn := func(resp http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(resp, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(resp, req)
	}

	return http.HandlerFunc(fn)
}
