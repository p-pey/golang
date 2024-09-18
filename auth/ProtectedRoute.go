package auth

import (
	"net/http"
)

func ProtectedRoute(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		bearer := request.Header.Get("bearer")

		if bearer == "" {
			http.Error(writer, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(writer, request)
	})
}
