package controllers

import (
	"io"
	"net/http"
)

type Controller struct{}

func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "test")
}
