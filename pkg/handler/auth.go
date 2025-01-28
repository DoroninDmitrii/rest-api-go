package handler

import (
	"net/http"
	restapitodo "rest-api-todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input restapitodo.User

	if error := c.BindJSON(&input); error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding: "required"`
	Password string `json:"password" binding: "required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if error := c.BindJSON(&input); error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
