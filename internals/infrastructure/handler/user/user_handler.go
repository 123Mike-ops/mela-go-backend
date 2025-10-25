package user

import (
	"auth-sso/internals/application/user"
	"auth-sso/utils/model"
	"fmt"
	"net/http"
	"strconv"

	RegisterUser "auth-sso/internals/infrastructure/handler"
	response "auth-sso/utils"

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
        err:=&model.ErrorResponse{
            Code:        http.StatusInternalServerError,
            Message:     "Failed to retrieve user",
            Description: err.Error(),
        }
        response.ErrorResponse(c,err )
        return
    }

    if user == nil {
        err:=&model.ErrorResponse{
            Code:        http.StatusNotFound,
            Message:     "User not found",
            Description: "No user exists with the given ID",
        }
        response.ErrorResponse(c,err )
        return
    }

    response.SuccessResponse(c, http.StatusOK, user, nil)

  
}
func (h *Handler) CreateUser(c *gin.Context) {
    var user RegisterUser.RegisterUser

    // Parse JSON body
    if err := c.ShouldBindJSON(&user); err != nil {
        errResp := &model.ErrorResponse{
            Code:        http.StatusBadRequest,
            Message:     "Invalid request body",
            Description: err.Error(),
        }
        response.ErrorResponse(c, errResp)
        return
    }

    // Call service layer to create the user
    createdUser, err := h.Service.CreateUser(c.Request.Context(), &user)
    if err != nil {
        errResp := &model.ErrorResponse{
            Code:        http.StatusInternalServerError,
            Message:     "Failed to create user",
            Description: err.Error(),
        }
        response.ErrorResponse(c, errResp)
        return
    }

    // Return success response
    response.SuccessResponse(c, http.StatusCreated, createdUser, nil)
}
                                                                    