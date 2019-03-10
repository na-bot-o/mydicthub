package auth

import (
	"net/http"

	"github.com/na-bot-o/mydicthub/util"
)

type Auth struct{}

func redirectToLoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "/login")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func (a *Auth) MustAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// cookie, err := r.Cookie("auth")
		// fmt.Println(err)
		// if err == http.ErrNoCookie || cookie.Value == "" {
		if util.IsCookie(r, "auth") {
			handler(w, r)
		} else {
			redirectToLoginHandler(w, r)
		}
	}
}

// func (a *Auth)
