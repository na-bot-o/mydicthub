package util

import (
	"net/http"
)

func IsCookie(r *http.Request, cookieName string) bool {
	cookie, err := r.Cookie(cookieName)
	if err == http.ErrNoCookie || cookie.Value == "" {
		return false
	}
	return true
}
