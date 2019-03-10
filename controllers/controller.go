package controllers

import (
	"fmt"
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
	fmt.Println(data)

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

func (c *Controller) CallbackHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		log.Fatal(err)
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
