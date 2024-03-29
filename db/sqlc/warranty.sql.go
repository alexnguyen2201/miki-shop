// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: warranty.sql

package db

import (
	"context"
	"database/sql"
)

const getWarranty = `-- name: GetWarranty :many
SELECT w.id, w.type, w.title, w.duration, w.price, w.times
FROM mt_product_type_warranty
JOIN warranties w on mt_product_type_warranty.warranty_id = w.id
where product_type_id = $1
`

type GetWarrantyRow struct {
	ID       int64         `json:"id"`
	Type     string        `json:"type"`
	Title    string        `json:"title"`
	Duration sql.NullInt64 `json:"duration"`
	Price    sql.NullInt64 `json:"price"`
	Times    sql.NullInt64 `json:"times"`
}

func (q *Queries) GetWarranty(ctx context.Context, productTypeID int64) ([]GetWarrantyRow, error) {
	rows, err := q.db.QueryContext(ctx, getWarranty, productTypeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetWarrantyRow{}
	for rows.Next() {
		var i GetWarrantyRow
		if err := rows.Scan(
			&i.ID,
			&i.Type,
			&i.Title,
			&i.Duration,
			&i.Price,
			&i.Times,
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
