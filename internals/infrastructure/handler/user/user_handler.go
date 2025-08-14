package user

import (
	"auth-sso/internals/application/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
    Service *user.Service
}

func NewHandler(service *user.Service) *Handler {
    return &Handler{Service: service}
}

func (h *Handler) GetUserByID(c *gin.Context) {
    idStr := c.Param("id");
    id, err := strconv.Atoi(idStr);
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
        return
    }

    user, err := h.Service.GetUser(c.Request.Context(), id)
    if err != nil {
        fmt.Print("error",err)
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}
