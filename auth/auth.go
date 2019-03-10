package auth

import (
	"fmt"
	"net/http"
)

type Auth struct{}

func redirectToLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (a *Auth) MustAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("auth")
		fmt.Println(err)
		if err == http.ErrNoCookie || cookie.Value == "" {
			redirectToLoginHandler(w, r)
		} else {
			handler(w, r)
		}
	}
}
