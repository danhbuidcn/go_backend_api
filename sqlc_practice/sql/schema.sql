-- Active: 1728134386556@@127.0.0.1@3308@shopdevgo
CREATE TABLE IF NOT EXISTS `products` (
productId int PRIMARY KEY AUTO_INCREMENT,
shopId int NOT NULL,
FOREIGN KEY (shopId) REFERENCES shops (shopId),
productName VARCHAR(255) NOT NULL,
productImage,
productDescription VARCHAR(255) NOT NULL,
);
-- Tạo thêm một schema shops
CREATE TABLE IF NOT EXISTS `shops`(
shopId int PRIMARY KEY AUTO_INCREMENT,
shopname VARCHAR(255) NOT NULL
);
