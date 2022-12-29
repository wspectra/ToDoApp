package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/api_server/internal/pkg/utils"
	"github.com/wspectra/api_server/internal/structure"
	"net/http"
	"strconv"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

	input := structure.List{}
	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.services.CreateList(id.(int), input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	utils.HttpResponseWriter(c.Writer, "list created successfully", http.StatusOK)
}

func (h *Handler) getAllLists(c *gin.Context) {
	id, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

	lists, err := h.services.GetLists(id.(int))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	allListsResponse := struct {
		Message []structure.List
	}{lists}
	c.JSON(http.StatusOK, allListsResponse)
}

func (h *Handler) getListById(c *gin.Context) {
	userId, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	lists, err := h.services.GetListById(userId.(int), listId)
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	ListsResponse := struct {
		Message structure.List
	}{lists}
	c.JSON(http.StatusOK, ListsResponse)
}

func (h *Handler) updateList(c *gin.Context) {
	_, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

	input := structure.UpdateListInput{}
	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	if err := h.services.UpdateList(listId, input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}
	utils.HttpResponseWriter(c.Writer, "list updated successfully", http.StatusOK)
}

func (h *Handler) deleteList(c *gin.Context) {
	_, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	err = h.services.DeleteList(listId)
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	utils.HttpResponseWriter(c.Writer, "list deleted successfully", http.StatusOK)

}
