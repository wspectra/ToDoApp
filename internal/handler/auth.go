package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/wspectra/ToDoApp/internal/pkg/utils"
	"github.com/wspectra/ToDoApp/internal/structure"
	"net/http"
)

// @Summary      SignUp
// @Tags         Auth
// @Description  create account
// @Accept       json
// @Produce      json
// @Param        input body structure.User true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	input := structure.User{}

	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.services.Authorization.AddNewUser(input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponseWriter(c.Writer, "new user created successfully", http.StatusOK)
}

// @Summary      SignIn
// @Tags         Auth
// @Description  create account
// @Accept       json
// @Produce      json
// @Param        input body structure.SignInUser true "account info"
// @Success      200  {object} utils.ResponseStruct
// @Failure      400  {object}  utils.ResponseStruct
// @Failure      500  {object}  utils.ResponseStruct
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	input := structure.SignInUser{}

	if err := c.BindJSON(&input); err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.services.Authorization.GetToken(input)
	if err != nil {
		utils.HttpResponseWriter(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.HttpResponseWriter(c.Writer, token, http.StatusOK)
}
