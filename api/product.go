package api

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/nguyenvanson2201/miki-shop/db/sqlc"
)

type createProductRequest struct {
	ProductTypeID int64          `json:"product_type_id" binding:"required"`
	Price         int64          `json:"price" binding:"required"`
	Size          sql.NullInt64  `json:"size"`
	Color         sql.NullString `json:"color"`
	Stock         int64          `json:"stock" binding:"required"`
}

func (server *Server) createProduct(ctx *gin.Context) {
	var req createProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		ProductTypeID: req.ProductTypeID,
		Price:         req.Price,
		Size:          req.Size,
		Color:         req.Color,
		Stock:         req.Stock,
	}

	product, err := server.store.CreateProduct(ctx, arg)
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

	ctx.JSON(http.StatusOK, product)
}

type getProductRequest struct {
	ProductTypeID int64 `uri:"id"  binding:"required,min=1"`
}

type policies struct {
	WarrantyContents []string `json:"warranty_contents"`
	Times            []string `json:"times"`
	PaidWarranties   []string `json:"paid_warranties"`
	Expenses         []string `json:"expenses"`
}

type getProductResponse struct {
	Name            string   `json:"name"`
	Images          []string `json:"images"`
	Rating          int64    `json:"rating"`
	QuantitySold    int64    `json:"quantity_sold"`
	Status          int64    `json:"status"`
	Price           []int64  `json:"price"`
	Sale            int64    `json:"sale"`
	Color           []string `json:"colors"`
	Descriptions    string   `json:"descriptions"`
	SalientFeatures []string `json:"salient_features"`
	Policies        policies `json:"policies"`
}

func (server *Server) getProduct(ctx *gin.Context) {
	var req getProductRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// productType, err := server.store.GetProduct(ctx, req.ProductTypeID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	// product, err := server.store.GetProduct(ctx, req.ID)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		ctx.JSON(http.StatusNotFound, errorResponse(err))
	// 		return
	// 	}
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }

	warranties, err := server.store.GetWarranty(ctx, req.ProductTypeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	var warrantyContents []string
	for _, w := range warranties {
		if w.Type == "free" {
			warrantyContents = append(warrantyContents, w.Title)
		}
	}

	var times []string
	for _, w := range warranties {
		if w.Type == "free" {
			if w.Duration.Int64 == 36500 {
				times = append(times, "Trọn đời")
			} else {
				timesStr := strconv.Itoa(int(w.Times.Int64))
				if w.Times.Int64 < 10 {
					timesStr = "0" + timesStr
				}
				times = append(times, timesStr+" lần")
			}
		}
	}

	var paidWarranties []string
	for _, w := range warranties {
		if w.Type == "paid" {
			paidWarranties = append(paidWarranties, w.Title)
		}
	}

	var expenses []string
	for _, w := range warranties {
		if w.Type == "paid" && w.Price.Valid {
			expenses = append(expenses, strconv.Itoa(int(w.Price.Int64))+"đ")
		}
	}

	policiesResponse := policies{
		WarrantyContents: warrantyContents,
		Times:            times,
		PaidWarranties:   paidWarranties,
		Expenses:         expenses,
	}

	ctx.JSON(http.StatusOK, policiesResponse)
}
