CREATE TABLE
    IF NOT EXISTS `order_items` (
        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `order_id` INT UNSIGNED NOT NULL,
        `product_id` INT UNSIGNED NOT NULL,
        `quantity` INT NOT NULL,
        `price` DECIMAL(10, 2) NOT NULL,
        PRIMARY KEY (`id`),
        CONSTRAINT `fk_order_items_order_id` FOREIGN KEY (`order_id`) REFERENCES orders (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT `fk_order_items_product_id` FOREIGN KEY (`product_id`) REFERENCES products (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    );