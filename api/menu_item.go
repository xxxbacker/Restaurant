package api

import (
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"net/http"
)

type getMenuItemRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

/*func (server *Server) getMenuItemForAdmin(ctx *gin.Context) {
	var req getMenuItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	menuItem, err := server.store.GetMenu_item(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menuItem)
} */

type createMenuItemRequest struct {
	Title    string `json:"title" binding:"required"`
	Category string `json:"category" binding:"required"`
	Price    int32  `json:"price" binding:"required"`
}

func (server *Server) createMenuItem(ctx *gin.Context) {
	var req createMenuItemRequest

	arg := db.CreateMenu_itemParams{
		Title:    req.Title,
		Category: req.Category,
		Price:    req.Price,
	}

	menuItem, err := server.store.CreateMenu_item(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menuItem)

	ctx.Set(menuItemIdCtx, menuItem.MenuID)
	ctx.Next()
}

type listMenuItemRequest struct {
	PageId   int32 `form:"Page_id" binding:"required,min=1"`
	PageSize int32 `form:"Page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listMenuItem(ctx *gin.Context) {
	var req listMenuItemRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.ListMenu_itemParams{
		Limit:  req.PageId,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	menuItem, err := server.store.ListMenu_item(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, menuItem)
}

type deleteMenuItemRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) DeleteMenuItem(ctx *gin.Context) {
	var req deleteMenuItemRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteMenu_item(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, true)
}
