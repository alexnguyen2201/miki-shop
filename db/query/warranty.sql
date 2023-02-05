-- name: GetWarranty :many
SELECT w.id, w.type, w.title, w.duration, w.price, w.times
FROM mt_product_type_warranty
JOIN warranties w on mt_product_type_warranty.warranty_id = w.id
where product_type_id = $1;