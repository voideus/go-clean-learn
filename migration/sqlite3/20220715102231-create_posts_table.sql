
-- +migrate Up
CREATE TABLE IF NOT EXISTS `posts` (
    `id` INT UNSIGNED NOT NULL,
    `title` VARCHAR(150) NOT NULL,
    `description` VARCHAR(400) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
);

-- +migrate Down
DROP TABLE IF EXISTS `posts`;
