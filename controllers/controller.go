package controllers

import (
	"io"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Controller struct{}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("auth")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	io.WriteString(w, "test")
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	err := t.ExecuteTemplate(w, "login.html", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func (c *Controller) AuthHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":

	}
}
