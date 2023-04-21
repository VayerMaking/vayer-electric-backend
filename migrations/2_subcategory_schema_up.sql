CREATE TABLE subcategory (
  id SERIAL PRIMARY KEY,
  name varchar(255) NOT NULL,
  description varchar(255),
  created_at timestamp NOT NULL,
  category_id int NOT NULL REFERENCES category(id)
);