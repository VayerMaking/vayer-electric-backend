CREATE TABLE `subcategory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_at` datetime NOT NULL,
  `category_id` int(11) NOT NULL FOREIGN KEY (`category_id`) REFERENCES `category` (`id`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;