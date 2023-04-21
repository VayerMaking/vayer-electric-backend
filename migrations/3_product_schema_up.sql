CREATE TABLE product (
  id SERIAL PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255),
  created_at timestamp NOT NULL,
  subcategory_id int NOT NULL REFERENCES subcategory(id),
  price numeric(10,2) NOT NULL,
  current_inventory int NOT NULL,
  image_url varchar(255) NOT NULL,
  brand varchar(255) NOT NULL,
  sku varchar(255) NOT NULL
);
