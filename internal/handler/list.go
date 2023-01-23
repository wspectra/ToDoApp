package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/ToDoApp/internal/pkg/utils"
	"github.com/wspectra/ToDoApp/internal/structure"
	"net/http"
	"strconv"
)

// @Summary      createList
// @Security	 ApiKeyAuth
// @Tags         List
// @Description  create list
// @Accept       json
// @Produce      json
// @Param        input body structure.List true "list info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list [post]
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

// @Summary      getAllLists
// @Security	 ApiKeyAuth
// @Tags         List
// @Description  returns all lists of the list
// @Produce      json
// @Success      200  {object} structure.AllListResponse
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list [get]
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
	allListsResponse := structure.AllListResponse{
		Status:  "Success",
		Message: lists}
	c.JSON(http.StatusOK, allListsResponse)
}

// @Summary      getListById
// @Security	 ApiKeyAuth
// @Tags         List
// @Description  return information about list
// @Produce      json
// @Success      200  {object} structure.ListResponse
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id [get]
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

	ListsResponse := structure.ListResponse{
		Status:  "Success",
		Message: lists}
	c.JSON(http.StatusOK, ListsResponse)
}

// @Summary      updateList
// @Security	 ApiKeyAuth
// @Tags         List
// @Description  update list
// @Accept       json
// @Produce      json
// @Param        input body structure.UpdateListInput true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id [put]
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

// @Summary      deleteList
// @Security	 ApiKeyAuth
// @Tags         List
// @Description  delete list
// @Accept       json
// @Produce      json
// @Param        input body structure.User true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      401  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /api/list/:id [delete]
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
