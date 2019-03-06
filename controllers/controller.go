package controllers

import (
	"log"
	"net/http"
	"text/template"

	"github.com/markbates/goth/gothic"
	"github.com/stretchr/objx"
)

type dictionary struct {
	userid string
	words  map[string]string
}

type Controller struct{}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})
	authcookie, _ := r.Cookie("auth")

	data["userdata"] = objx.MustFromBase64(authcookie.Value)

	t := template.Must(template.ParseFiles("templates/index.html"))
	err := t.ExecuteTemplate(w, "index.html", data)

	if err != nil {
		log.Fatal(err)
	}
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	err := t.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func (c *Controller) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Errorf(err)
		return
	}
	authCookieValue := objx.New(map[string]interface{}{
		"name": user.Name,
		"id":   user.UserID,
	}).MustBase64()

	http.SetCookie(w, &http.Cookie{
		Name:  "auth",
		Value: authCookieValue,
		Path:  "/",
	})

	w.Header()["location"] = []string{"/"}
	w.WriteHeader(http.StatusTemporaryRedirect)

}
