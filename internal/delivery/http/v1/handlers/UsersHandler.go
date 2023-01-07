package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stickpro/vpn-sass/internal/delivery/resource"
	"net/http"
)

func (h *Handler) UsersPageIndex(c *gin.Context) {
	users, err := h.services.Users.LoadAll()
	if err != nil {
		resource.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("", h.UsersPageIndex)
	}
}
