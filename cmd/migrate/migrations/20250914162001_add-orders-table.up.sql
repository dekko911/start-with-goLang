CREATE TABLE
    IF NOT EXISTS `orders` (
        `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
        `userId` INT UNSIGNED NOT NULL,
        `total` DECIMAL(10, 2) NOT NULL,
        `status` ENUM ('pending', 'completed', 'canceled') NOT NULL DEFAULT 'pending',
        `address` TEXT NOT NULL,
        `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        CONSTRAINT `fk_orders_userId` FOREIGN KEY (`userId`) REFERENCES users (`id`) ON DELETE CASCADE ON UPDATE CASCADE
    );