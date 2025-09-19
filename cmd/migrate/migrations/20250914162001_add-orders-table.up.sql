CREATE TABLE
    IF NOT EXISTS `orders` (
        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `userID` INT UNSIGNED NOT NULL,
        `total` DECIMAL(10, 2) NOT NULL,
        `status` ENUM ('pending', 'completed', 'canceled') NOT NULL DEFAULT 'pending',
        `address` TEXT NOT NULL,
        `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        CONSTRAINT `fk_orders_userID` FOREIGN KEY (`userID`) REFERENCES users (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    );