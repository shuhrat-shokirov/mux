package unauthenticated

import (
	"context"
	"net/http"
)

func UnAuthenticated(vasya func(ctx context.Context) bool, redirect bool, redirectURL string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(writer http.ResponseWriter, request *http.Request) {
			if !vasya(request.Context()) {
				next(writer, request)
				return
			}

			if redirect {
				http.Redirect(writer, request, redirectURL, http.StatusTemporaryRedirect)
				return
			}

			http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}
}
