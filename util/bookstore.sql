DROP DATABASE IF EXISTS bookstore;
CREATE DATABASE bookstore;
USE bookstore;

DROP TABLE IF EXISTS `books`;
CREATE TABLE `books` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  `author` varchar(50) NOT NULL,
  `published_date` date NOT NULL,
  `isbn` varchar(50) NOT NULL,
  `number_of_pages` int NOT NULL,
  `cover_image` json NOT NULL,
  `language` varchar(20) NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;