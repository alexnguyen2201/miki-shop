package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/nguyenvanson2201/miki-shop/db/sqlc"
)

type createProductTypeRequest struct {
	Title           string   `json:"title" binding:"required"`
	SalientFeatures []string `json:"salient_features" binding:"required"`
	Descriptions    []string `json:"descriptions" binding:"required"`
}

func (server *Server) createProductType(ctx *gin.Context) {
	var req createProductTypeRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductTypeParams{
		Title:           req.Title,
		SalientFeatures: req.SalientFeatures,
		Descriptions:    req.Descriptions,
	}

	productType, err := server.store.CreateProductType(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productType)
}

type getProductTypeRequest struct {
	ID int64 `uri:"id"  binding:"required,min=1"`
}

func (server *Server) getProductType(ctx *gin.Context) {
	var req getProductTypeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productType, err := server.store.GetProductType(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productType)

}

type getProductTypeListRequest struct {
	PageID   int32 `form:"page_id,default=1"  binding:"required,min=1"`
	PageSize int32 `form:"page_size,default=5"  binding:"required,min=5,max=10"`
}

func (server *Server) getProductTypeList(ctx *gin.Context) {
	var req getProductTypeListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetListProductTypesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	productTypes, err := server.store.GetListProductTypes(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, productTypes)

}
