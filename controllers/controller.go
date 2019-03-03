package controllers

import (
	"io"
	"net/http"
)

type Controller struct{}

func (c *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("session")

	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	io.WriteString(w, "test")
}

func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "this is login page")
}
