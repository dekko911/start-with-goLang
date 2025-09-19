CREATE TABLE
    IF NOT EXISTS `order_items` (
        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `orderID` INT UNSIGNED NOT NULL,
        `productID` INT UNSIGNED NOT NULL,
        `quantity` INT NOT NULL,
        `price` DECIMAL(10, 2) NOT NULL,
        PRIMARY KEY (`id`),
        CONSTRAINT `fk_order_items_orderID` FOREIGN KEY (`orderID`) REFERENCES orders (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
        CONSTRAINT `fk_order_items_productID` FOREIGN KEY (`productID`) REFERENCES products (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    );