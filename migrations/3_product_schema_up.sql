CREATE TABLE `product` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_at` datetime NOT NULL,
  `subcategory_id` int(11) NOT NULL FOREIGN KEY (`subcategory_id`) REFERENCES `subcategory` (`id`),
  `price` decimal(10,2) NOT NULL,
  `currentInventory` int(11) NOT NULL,
  `image` varchar(255) NOT NULL,
  `brand` varchar(255) NOT NULL,
  `sku` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;