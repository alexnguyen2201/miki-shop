// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: product_type.sql

package db

import (
	"context"

	"github.com/lib/pq"
)

const createProductType = `-- name: CreateProductType :one
INSERT INTO product_types (
  title,
  salient_features,
  descriptions
) VALUES (
  $1, $2, $3
) RETURNING id, title, salient_features, descriptions, created_at, updated_at, created_by, updated_by, discount
`

type CreateProductTypeParams struct {
	Title           string   `json:"title"`
	SalientFeatures []string `json:"salient_features"`
	Descriptions    []string `json:"descriptions"`
}

func (q *Queries) CreateProductType(ctx context.Context, arg CreateProductTypeParams) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, createProductType, arg.Title, pq.Array(arg.SalientFeatures), pq.Array(arg.Descriptions))
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Title,
		pq.Array(&i.SalientFeatures),
		pq.Array(&i.Descriptions),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Discount,
	)
	return i, err
}

const deleteProductType = `-- name: DeleteProductType :exec
DELETE FROM product_types WHERE id = $1
`

func (q *Queries) DeleteProductType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteProductType, id)
	return err
}

const getListProductTypes = `-- name: GetListProductTypes :many
SELECT id, title, salient_features, descriptions, created_at, updated_at, created_by, updated_by, discount FROM product_types
ORDER BY id
LIMIT $1
OFFSET $2
`

type GetListProductTypesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetListProductTypes(ctx context.Context, arg GetListProductTypesParams) ([]ProductType, error) {
	rows, err := q.db.QueryContext(ctx, getListProductTypes, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ProductType{}
	for rows.Next() {
		var i ProductType
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			pq.Array(&i.SalientFeatures),
			pq.Array(&i.Descriptions),
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatedBy,
			&i.UpdatedBy,
			&i.Discount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductType = `-- name: GetProductType :one
SELECT id, title, salient_features, descriptions, created_at, updated_at, created_by, updated_by, discount FROM product_types
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetProductType(ctx context.Context, id int64) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, getProductType, id)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Title,
		pq.Array(&i.SalientFeatures),
		pq.Array(&i.Descriptions),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Discount,
	)
	return i, err
}

const getProductTypeForUpdate = `-- name: GetProductTypeForUpdate :one
SELECT id, title, salient_features, descriptions, created_at, updated_at, created_by, updated_by, discount FROM product_types
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductTypeForUpdate(ctx context.Context, id int64) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, getProductTypeForUpdate, id)
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Title,
		pq.Array(&i.SalientFeatures),
		pq.Array(&i.Descriptions),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Discount,
	)
	return i, err
}

const updateProductType = `-- name: UpdateProductType :one
UPDATE product_types
SET salient_features = COALESCE($2, salient_features), 
    descriptions = COALESCE($3, descriptions)
WHERE id = $1
RETURNING id, title, salient_features, descriptions, created_at, updated_at, created_by, updated_by, discount
`

type UpdateProductTypeParams struct {
	ID              int64    `json:"id"`
	SalientFeatures []string `json:"salient_features"`
	Descriptions    []string `json:"descriptions"`
}

func (q *Queries) UpdateProductType(ctx context.Context, arg UpdateProductTypeParams) (ProductType, error) {
	row := q.db.QueryRowContext(ctx, updateProductType, arg.ID, pq.Array(arg.SalientFeatures), pq.Array(arg.Descriptions))
	var i ProductType
	err := row.Scan(
		&i.ID,
		&i.Title,
		pq.Array(&i.SalientFeatures),
		pq.Array(&i.Descriptions),
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatedBy,
		&i.UpdatedBy,
		&i.Discount,
	)
	return i, err
}