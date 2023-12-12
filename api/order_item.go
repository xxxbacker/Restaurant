package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"net/http"
)

type getOrderItemRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getOrderItem(ctx *gin.Context) {
	var req getOrderItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	orderItem, err := server.store.GetOrder_item(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderItem)
}

type createOrderItemRequest struct {
	Title  string `json:"title" binding:"required"`
	Price  int32  `json:"price" binding:"required"`
	MenuID int32  `json:"menu_id" binding:"required"`
	OrdID  int32  `json:"ord_id" binding:"required"`
}

func (server *Server) createOrderItem(ctx *gin.Context) {
	var req createOrderItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	menu, val := server.validMenu(ctx, req.MenuID)
	if !val {
		return
	}

	ord, valid := server.validOrd(ctx, req.OrdID)
	if !valid {
		return
	}
	arg := db.CreateOrder_itemParams{
		Title: sql.NullString{
			String: req.Title,
			Valid:  true,
		},
		Price: req.Price,
		MenuID: sql.NullInt64{
			Int64: int64(menu.MenuID),
			Valid: val,
		},
		OrdID: sql.NullInt64{
			Int64: int64(ord.OrdID),
			Valid: valid,
		},
	}

	orderItem, err := server.store.CreateOrder_item(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, orderItem)
}

func (server *Server) validMenu(ctx *gin.Context, menuId int32) (db.MenuItem, bool) {
	menu, err := server.store.GetMenu_item(ctx, menuId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return menu, false
	}

	return menu, true
}

type listMenuRequest struct {
	PageId   int32 `form:"Page_id" binding:"required,min=1"`
	PageSize int32 `form:"Page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMenu(ctx *gin.Context) {
	var req listMenuRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.ListMenu_itemParams{
		Limit:  req.PageId,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	menuItems, err := server.store.ListMenu_item(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menuItems)
}
