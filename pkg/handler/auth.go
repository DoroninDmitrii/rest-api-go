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

func (h *Handler) signIn(c *gin.Context) {

}
