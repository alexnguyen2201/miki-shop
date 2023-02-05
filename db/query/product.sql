-- name: CreateProduct :one
INSERT INTO products (
  product_type_id,
  price,
  size,
  color,
  stock
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1 LIMIT 1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :one
UPDATE products
SET price = $2
WHERE id = $1
RETURNING *;

-- name: DeleteProductB :exec
DELETE FROM products WHERE id = $1;