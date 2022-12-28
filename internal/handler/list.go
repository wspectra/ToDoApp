package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/api_server/internal/pkg/utils"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	id, ok := c.Get(userCt)
	if !ok {
		utils.HttpResponseWriter(c.Writer, "user id not found", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
