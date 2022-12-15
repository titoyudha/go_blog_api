package controllers

import (
	"net/http"

	"github.com/titoyudha/go_blog_api/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome to this API")
}
