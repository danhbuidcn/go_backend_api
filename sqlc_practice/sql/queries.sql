-- name: CreateShop :exec
INSERT INTO shops (shopname) VALUES(?);

-- name: GetShop :one
SELECT * FROM shops WHERE shopId = ?;

-- name: GetShops :many
SELECT * FROM shops;

-- name: CreateProduct :exec
INSERT INTO products (shopId, productName) VALUES (?, ?);

-- name: GetProduct :one
SELECT * FROM products WHERE productId = ?;

-- name: GetProducts :many
SELECT * FROM products;