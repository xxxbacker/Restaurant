package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"net/http"
)

type getChequeRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCheque(ctx *gin.Context) {
	var req getChequeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	cheque, err := server.store.GetCheque(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cheque)
}

type createChequeRequest struct {
	Price int32 `json:"price" binding:"required"`
	OrdId int32 `json:"ord_id" binding:"required"`
}

func (server *Server) createCheque(ctx *gin.Context) {
	var req createChequeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ord, val := server.validOrd(ctx, req.OrdId)
	if !val {
		return
	}

	arg := db.CreateChequeParams{
		Price: req.Price,
		OrdID: sql.NullInt64{
			Int64: int64(ord.OrdID),
			Valid: val,
		},
	}

	cheque, err := server.store.CreateCheque(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cheque)
}

func (server *Server) validOrd(ctx *gin.Context, ordId int32) (db.Ord, bool) {
	ord, err := server.store.GetOrd(ctx, ordId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return ord, false
	}

	return ord, true
}

type listChequeRequest struct {
	Limit int32 `form:"limit" binding:"required,min=1"`
}

func (server *Server) listCheque(ctx *gin.Context) {
	var req listChequeRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	cheques, err := server.store.ListCheque(ctx, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, cheques)
}
