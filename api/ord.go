package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"golangRestaurantManagement/utils"
	"net/http"
	"time"
)

const (
	courierIdCtx = "courierId"
	ordIdCtx     = "ordId"
)

type getOrdForAdminRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getOrdForAdmin(ctx *gin.Context) {
	var req getOrdForAdminRequest
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

func (server *Server) getOrdForUser(ctx *gin.Context) {
	ordId, err := getOrdId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ord, err := server.store.GetOrd(ctx, ordId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ord)
}

func (server *Server) createOrd(ctx *gin.Context) {
	accountId, err := getAccountId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	account, err := server.store.GetAccount(ctx, int32(accountId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	n := utils.RandomModuloTen()
	courier, err := server.store.GetCourier(ctx, n)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Set(courierIdCtx, courier.CourierID)
	ctx.Next()

	arg := db.CreateOrdParams{
		AccountID: sql.NullInt64{
			Int64: int64(account.AccountID),
			Valid: true,
		},
		CourierID: sql.NullInt64{
			Int64: int64(courier.CourierID),
			Valid: true,
		},
		CreatedAt: time.Now(),
	}

	ord, err := server.store.CreateOrd(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, ord)

	ctx.Set(ordIdCtx, ord.OrdID)
	ctx.Next()
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
