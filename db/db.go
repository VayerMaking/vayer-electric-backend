package db

import (
	"database/sql"
	"fmt"
	"time"

	"vayer-electric-backend/env"
	"vayer-electric-backend/logging"
	"vayer-electric-backend/structs"

	"github.com/DavidHuie/gomigrate"
	_ "github.com/lib/pq"
)

type DbSource struct {
	conn *sql.DB
}

var log = logging.GetLogger()

func CreateDbSource(dsn string) (DbSource, error) {
	d, err := sql.Open("postgres", dsn)

	if err != nil {
		return DbSource{}, err
	}

	go func() {
		for {
			d.Ping()
			time.Sleep(time.Second * 30)
		}
	}()

	d.SetMaxOpenConns(6)
	d.SetMaxIdleConns(2)

	return DbSource{
		conn: d,
	}, nil
}

func GetDbSource() DbSource {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
	)
	src, err := CreateDbSource(dsn)

	if err != nil {
		panic(err)
	}

	return src
}

func (s DbSource) ValidateConnection() bool {
	return s.conn.Ping() == nil
}

func (s DbSource) Migrate(path string) error {
	migrator, _ := gomigrate.NewMigrator(s.conn, gomigrate.Postgres{}, path)
	defer s.conn.Close()
	return migrator.Migrate()
}

func (s DbSource) InsertProduct(name string, description string, subcategory_id int, price float64, currentInventory int, imageUrl string, brand string, sku string) error {
	_, err := s.conn.Exec("INSERT INTO product (name, description, subcategory_id, price, current_inventory, image_url, brand, sku, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", name, description, subcategory_id, price, currentInventory, imageUrl, brand, sku, time.Now())
	defer s.conn.Close()
	return err
}

func (s DbSource) UpdateProduct(id int, name string, description string, subcategory_id int, price float64, currentInventory int, image string, brand string, sku string) error {
	_, err := s.conn.Exec("UPDATE product SET name = $1, description = $2, subcategory_id = $3, price = $4, current_inventory = $5, image = $6, brand = $7, sku = $8, updated_at = $9 WHERE id = $10", name, description, subcategory_id, price, currentInventory, image, brand, sku, time.Now(), id)
	defer s.conn.Close()
	return err
}

func (s DbSource) DeleteProduct(id int) error {
	_, err := s.conn.Exec("DELETE FROM product WHERE id = $1", id)
	defer s.conn.Close()
	return err
}

func (s DbSource) GetProducts() ([]structs.Product, error) {
	rows, err := s.conn.Query("SELECT * FROM product")

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	products := make([]structs.Product, 0)

	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return products, nil
}

func (s DbSource) GetProductById(id int) (structs.Product, error) {
	var product structs.Product
	err := s.conn.QueryRow("SELECT * FROM product WHERE id = $1", id).Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

	if err != nil {
		log.Error(err.Error())
		return structs.Product{}, err
	}

	defer s.conn.Close()

	return product, nil
}

func (s DbSource) GetProductByName(name string) (structs.Product, error) {
	var product structs.Product
	err := s.conn.QueryRow("SELECT * FROM product WHERE name = $1", name).Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

	if err != nil {
		log.Error(err.Error())
		return structs.Product{}, err
	}

	defer s.conn.Close()

	return product, nil

}

func (s DbSource) GetProductsBySubcategoryId(subcategory_id int) ([]structs.Product, error) {
	rows, err := s.conn.Query("SELECT * FROM product WHERE subcategory_id = $1", subcategory_id)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	products := make([]structs.Product, 0)

	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return products, nil
}

func (s DbSource) GetProductsByCategoryId(categoryId int) ([]structs.Product, error) {
	rows, err := s.conn.Query("SELECT * FROM product WHERE subcategory_id IN (SELECT id FROM subcategory WHERE category_id = $1)", categoryId)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	products := make([]structs.Product, 0)

	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return products, nil
}

func (s DbSource) GetProductsByCategoryName(categoryName string) ([]structs.Product, error) {
	rows, err := s.conn.Query("SELECT * FROM product WHERE subcategory_id IN (SELECT id FROM subcategory WHERE category_id = (SELECT id FROM category WHERE name = $1))", categoryName)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	products := make([]structs.Product, 0)

	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.CreatedAt, &product.SubcategoryId, &product.Price, &product.CurrentInventory, &product.ImageUrl, &product.Brand, &product.Sku)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		products = append(products, product)

	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return products, nil
}

func (s DbSource) InsertSubcategory(name string, description string, category_id int, image_url string) error {
	_, err := s.conn.Exec("INSERT INTO subcategory (name, description, category_id, created_at, image_url) VALUES ($1, $2, $3, $4, $5)", name, description, category_id, time.Now(), image_url)

	defer s.conn.Close()

	return err
}

func (s DbSource) UpdateSubcategory(id int, name string, description string, category_id int, image_url string) error {
	_, err := s.conn.Exec("UPDATE subcategory SET name = $1, description = $2, category_id = $3, updated_at = $4, image_url = $5 WHERE id = $6", name, description, category_id, time.Now(), image_url, id)

	defer s.conn.Close()

	return err
}

func (s DbSource) DeleteSubcategory(id int) error {
	_, err := s.conn.Exec("DELETE FROM subcategory WHERE id = $1", id)

	defer s.conn.Close()

	return err
}

func (s DbSource) GetSubcategories() ([]structs.Subcategory, error) {
	rows, err := s.conn.Query("SELECT * FROM subcategory")

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	subcategories := make([]structs.Subcategory, 0)

	for rows.Next() {
		var subcategory structs.Subcategory
		err := rows.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Description, &subcategory.CreatedAt, &subcategory.CategoryId, &subcategory.ImageUrl)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		subcategories = append(subcategories, subcategory)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return subcategories, nil
}

func (s DbSource) GetSubcategoryById(id int) (structs.Subcategory, error) {
	var subcategory structs.Subcategory
	err := s.conn.QueryRow("SELECT * FROM subcategory WHERE id = $1", id).Scan(&subcategory.Id, &subcategory.Name, &subcategory.Description, &subcategory.CreatedAt, &subcategory.CategoryId, &subcategory.ImageUrl)

	if err != nil {
		log.Error(err.Error())
		return structs.Subcategory{}, err
	}

	defer s.conn.Close()

	return subcategory, nil
}

func (s DbSource) GetSubcategoryByName(name string) (structs.Subcategory, error) {
	var subcategory structs.Subcategory
	err := s.conn.QueryRow("SELECT * FROM subcategory WHERE name = $1", name).Scan(&subcategory.Id, &subcategory.Name, &subcategory.Description, &subcategory.CreatedAt, &subcategory.CategoryId, &subcategory.ImageUrl)

	if err != nil {
		log.Error(err.Error())
		return structs.Subcategory{}, err
	}

	defer s.conn.Close()

	return subcategory, nil
}

func (s DbSource) InsertCategory(name string, description string, image_url string) error {
	_, err := s.conn.Exec("INSERT INTO category (name, description, created_at, image_url) VALUES ($1, $2, $3, $4)", name, description, time.Now(), image_url)
	defer s.conn.Close()

	return err
}

func (s DbSource) UpdateCategory(id int, name string, description string, image_url string) error {
	_, err := s.conn.Exec("UPDATE category SET name = $1, description = $2, updated_at = $3, image_url = $4 WHERE id = $5", name, description, time.Now(), image_url, id)
	defer s.conn.Close()

	return err
}

func (s DbSource) DeleteCategory(id int) error {
	_, err := s.conn.Exec("DELETE FROM category WHERE id = $1", id)
	defer s.conn.Close()

	return err
}

func (s DbSource) GetCategories() ([]structs.Category, error) {
	rows, err := s.conn.Query("SELECT * FROM category")

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	categories := make([]structs.Category, 0)

	for rows.Next() {
		var category structs.Category
		err := rows.Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt, &category.ImageUrl)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return categories, nil
}

func (s DbSource) GetCategoryById(id int) (structs.Category, error) {
	var category structs.Category
	err := s.conn.QueryRow("SELECT * FROM category WHERE id = $1", id).Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt, &category.ImageUrl)

	if err != nil {
		log.Error(err.Error())
		return structs.Category{}, err
	}

	defer s.conn.Close()

	return category, nil
}

func (s DbSource) GetCategoryByName(name string) (structs.Category, error) {
	var category structs.Category
	err := s.conn.QueryRow("SELECT * FROM category WHERE name = $1", name).Scan(&category.Id, &category.Name, &category.Description, &category.CreatedAt, &category.ImageUrl)

	if err != nil {
		log.Error(err.Error())
		return structs.Category{}, err
	}

	defer s.conn.Close()

	return category, nil
}

func (s DbSource) GetSubcategoriesByCategoryId(category_id int) ([]structs.Subcategory, error) {
	rows, err := s.conn.Query("SELECT * FROM subcategory WHERE category_id = $1", category_id)

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	subcategories := make([]structs.Subcategory, 0)

	for rows.Next() {
		var subcategory structs.Subcategory
		err := rows.Scan(&subcategory.Id, &subcategory.Name, &subcategory.Description, &subcategory.CreatedAt, &subcategory.CategoryId, &subcategory.ImageUrl)

		if err != nil {
			log.Error(err.Error())
			return nil, err
		}

		subcategories = append(subcategories, subcategory)
	}

	if err = rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}

	defer s.conn.Close()

	return subcategories, nil
}
