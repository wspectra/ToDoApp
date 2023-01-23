package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/ToDoApp/internal/pkg/utils"
	"github.com/wspectra/ToDoApp/internal/structure"
	"net/http"
	"strconv"
)

// @Summary      createItem
// @Security	 ApiKeyAuth
// @Tags         Item
// @Description  create item
// @Accept       json
// @Produce      json
// @Param        input body structure.Item true "item info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id/items [post]
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

// @Summary      getAllItems
// @Security	 ApiKeyAuth
// @Tags         Item
// @Description  returns all items of the list
// @Produce      json
// @Success      200  {object} structure.AllItemResponse
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id/items [get]
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

	ListsResponse := structure.AllItemResponse{
		Status:  "success",
		Message: output,
	}
	c.JSON(http.StatusOK, ListsResponse)
}

// @Summary      getItemById
// @Security	 ApiKeyAuth
// @Tags         Item
// @Description  return information about item
// @Produce      json
// @Success      200  {object} structure.ItemResponse
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id/items/:item_id [get]
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

	ListsResponse := structure.ItemResponse{
		Status:  "success",
		Message: output}
	c.JSON(http.StatusOK, ListsResponse)
}

// @Summary      updateItem
// @Security	 ApiKeyAuth
// @Tags         Item
// @Description  update item
// @Accept       json
// @Produce      json
// @Param        input body structure.UpdateItemInput true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id/items/:item_id [put]
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

// @Summary      deleteItem
// @Security	 ApiKeyAuth
// @Tags         Item
// @Description  delete item
// @Accept       json
// @Produce      json
// @Param        input body structure.User true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id/items/:item_id [delete]
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
