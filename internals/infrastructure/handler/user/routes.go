package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(rg *gin.RouterGroup, handler *Handler) {
    users := rg.Group("/users")
    users.GET("/:id", handler.GetUserByID)
}
