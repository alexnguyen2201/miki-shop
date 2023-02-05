// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"context"
)

type Querier interface {
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateProductType(ctx context.Context, arg CreateProductTypeParams) (ProductType, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteProductB(ctx context.Context, id int64) error
	DeleteProductType(ctx context.Context, id int64) error
	GetListProductTypes(ctx context.Context, arg GetListProductTypesParams) ([]ProductType, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductForUpdate(ctx context.Context, id int64) (Product, error)
	GetProductType(ctx context.Context, id int64) (ProductType, error)
	GetProductTypeForUpdate(ctx context.Context, id int64) (ProductType, error)
	GetUserByID(ctx context.Context, id int64) (User, error)
	GetUserByUserName(ctx context.Context, username string) (User, error)
	GetWarranty(ctx context.Context, productTypeID int64) ([]GetWarrantyRow, error)
	ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error)
	UpdateProductType(ctx context.Context, arg UpdateProductTypeParams) (ProductType, error)
}

var _ Querier = (*Queries)(nil)
