package api

import (
	"github.com/gin-gonic/gin"
	db "golangRestaurantManagement/db/sqlc"
	"golangRestaurantManagement/utils"
	"net/http"
)

type getCourierRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) getCourier(ctx *gin.Context) {
	var req getCourierRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	courier, err := server.store.GetCourier(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, courier)
}

func (server *Server) createCourier(ctx *gin.Context) {
	req := utils.CreateRandomCourierRequest()

	arg := db.CreateCourierParams{
		Title: req.Title,
		Phone: req.Phone,
	}

	courier, err := server.store.CreateCourier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, courier)
}

type listCourierRequest struct {
	PageId   int32 `form:"Page_id" binding:"required,min=1"`
	PageSize int32 `form:"Page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listCourier(ctx *gin.Context) {
	var req listCourierRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := db.ListCourierParams{
		Limit:  req.PageId,
		Offset: (req.PageId - 1) * req.PageSize,
	}
	couriers, err := server.store.ListCourier(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, couriers)

}
