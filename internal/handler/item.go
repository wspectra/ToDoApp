package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/ToDoApp/internal/pkg/utils"
	"github.com/wspectra/ToDoApp/internal/structure"
	"net/http"
	"strconv"
)

func (h *Handler) createItem(c *gin.Context) {
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

	input := structure.Item{}

	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.services.Item.CreateItem(userId.(int), listId, input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponseWriter(c.Writer, "item created successfully", http.StatusOK)

}

func (h *Handler) getAllItems(c *gin.Context) {
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

	output, err := h.services.Item.GetItems(userId.(int), listId)
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ListsResponse := struct {
		status  string
		Message []structure.Item
	}{"success",
		output}
	c.JSON(http.StatusOK, ListsResponse)
}

func (h *Handler) getItemById(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	output, err := h.services.Item.GetItemById(userId.(int), listId, itemId)
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	ListsResponse := struct {
		status  string
		Message structure.Item
	}{"success",
		output}
	c.JSON(http.StatusOK, ListsResponse)
}

func (h *Handler) updateItem(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	input := structure.UpdateItemInput{}
	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := h.services.Item.UpdateItem(userId.(int), listId, itemId, input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.HttpResponseWriter(c.Writer, "item updated successfully", http.StatusOK)
}

func (h *Handler) deleteItem(c *gin.Context) {
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

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		utils.HttpResponseWriter(c.Writer, "invalid id param", http.StatusInternalServerError)
		return
	}

	if err := h.services.Item.DeleteItem(userId.(int), listId, itemId); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.HttpResponseWriter(c.Writer, "item deleted successfully", http.StatusOK)

}
