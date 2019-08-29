package controllers

import (
	"net/http"

	"github.com/victoroliveira1605/PostsApi/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
