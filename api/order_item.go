package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"net/http"
)

type createOrderItemRequest struct {
	Title string `json:"title" binding:"required"`
	Price int32  `json:"price" binding:"required"`
}

func (server *Server) createOrderItem(ctx *gin.Context) {
	var req createOrderItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	menuItemId, err := getMenuItemId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ordId, err := getOrdId(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateOrder_itemParams{
		Title: sql.NullString{
			String: req.Title,
			Valid:  true,
		},
		Price: req.Price,
		MenuID: sql.NullInt64{
			Int64: int64(menuItemId),
			Valid: true,
		},
		OrdID: sql.NullInt64{
			Int64: int64(ordId),
			Valid: true,
		},
	}

	orderItem, err := server.store.CreateOrder_item(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderItem)
}

type listOrderItemRequest struct {
	Limit int32 `form:"limit" binding:"required,min=1"`
}

func (server *Server) listOrderItem(ctx *gin.Context) {
	var req listOrderItemRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	orderItems, err := server.store.ListOrder_item(ctx, req.Limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderItems)
}
