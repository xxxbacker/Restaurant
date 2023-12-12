package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"net/http"
)

type getOrdRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getOrd(ctx *gin.Context) {
	var req getOrdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ord, err := server.store.GetOrd(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ord)
}

type createOrdRequest struct {
	AccountID int32 `json:"account_id" binding:"required"`
	CourierID int32 `json:"courier_id" binding:"required"`
}

func (server *Server) createOrd(ctx *gin.Context) {
	var req createOrdRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	account, val := server.validAccount(ctx, req.AccountID)
	if !val {
		return
	}

	courier, valid := server.validCourier(ctx, req.CourierID)
	if !valid {
		return
	}
	arg := db.CreateOrdParams{
		AccountID: sql.NullInt64{
			Int64: int64(account.AccountID),
			Valid: val,
		},
		CourierID: sql.NullInt64{
			Int64: int64(courier.CourierID),
			Valid: valid,
		},
	}

	ord, err := server.store.CreateOrd(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ord)
}

func (server *Server) validAccount(ctx *gin.Context, accountId int32) (db.Account, bool) {
	account, err := server.store.GetAccount(ctx, accountId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return account, false
	}

	return account, true
}

func (server *Server) validCourier(ctx *gin.Context, courierId int32) (db.Courier, bool) {
	courier, err := server.store.GetCourier(ctx, courierId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return courier, false
	}

	return courier, true
}

type listOrdRequest struct {
	Limit int32 `form:"limit" binding:"required,min=1"`
}

func (server *Server) listOrd(ctx *gin.Context) {
	var req listOrdRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ords, err := server.store.ListOrd(ctx, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ords)
}
