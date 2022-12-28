package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/api_server/internal/pkg/utils"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCt              = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		utils.HttpResponseWriter(c.Writer, "empty auth header", http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		utils.HttpResponseWriter(c.Writer, "invalid auth header", http.StatusUnauthorized)
		return
	}

	userId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusUnauthorized)
		return
	}
	c.Set(userCt, userId)
}
