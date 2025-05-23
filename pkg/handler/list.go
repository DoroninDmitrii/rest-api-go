package handler

import (
	"net/http"
	restapitodo "rest-api-todo"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)

	if err != nil {
		return
	}

	var input restapitodo.TodoList
	if error := c.BindJSON(&input); error != nil {
		newErrorResponse(c, http.StatusBadRequest, error.Error())
	}

	id, err := h.services.TodoList.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
