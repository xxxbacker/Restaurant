package api

import (
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"golangRestaurantManagement/utils"
	"net/http"
	"time"
)

type createAccountRequest struct {
	Post      string `json:"post" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
}

func (server *Server) signUp(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Post:      "user",
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  utils.GeneratePasswordHash(req.Password),
		Email:     req.Email,
		Phone:     req.Phone,
		CreatedAt: time.Now(),
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		if arg.Post != "" {
			ctx.Abort()
		}
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type signInRequest struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) signIn(ctx *gin.Context) {
	var req signInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	token, err := server.GenerateToken(ctx, req.Phone, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, token)
}
