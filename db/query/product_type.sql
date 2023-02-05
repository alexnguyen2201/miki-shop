-- name: CreateProductType :one
INSERT INTO product_types (
  title,
  salient_features,
  descriptions
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetProductType :one
SELECT * FROM product_types
WHERE id = $1 LIMIT 1;

-- name: GetProductTypeForUpdate :one
SELECT * FROM product_types
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetListProductTypes :many
SELECT * FROM product_types
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProductType :one
UPDATE product_types
SET salient_features = COALESCE($2, salient_features), 
    descriptions = COALESCE($3, descriptions)
WHERE id = $1
RETURNING *;

-- name: DeleteProductType :exec
DELETE FROM product_types WHERE id = $1;