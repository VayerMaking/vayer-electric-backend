package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"vayer-electric-backend/db"

	"github.com/go-chi/chi/v5"
)

func GetCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs := db.GetDbSource()
		categories, err := dbs.GetCategories()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(categories)
	}
}

func GetCategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		parsedId, err := strconv.Atoi(id)
		dbs := db.GetDbSource()
		category, err := dbs.GetCategoryById(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(category)
	}
}

func GetCategoryByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		dbs := db.GetDbSource()
		category, err := dbs.GetCategoryByName(name)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(category)
	}
}

func CreateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)

		name := body.Name
		description := body.Description

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		err = dbs.InsertCategory(name, description)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Id = strings.TrimSpace(body.Id)
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)

		id := body.Id
		name := body.Name
		description := body.Description

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		err = dbs.UpdateCategory(parsedId, name, description)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteCategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		err = dbs.DeleteCategory(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetSubcategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs := db.GetDbSource()
		subcategories, err := dbs.GetSubcategories()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(subcategories)
	}
}

func GetSubcategoryById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		subcategory, err := dbs.GetSubcategoryById(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(subcategory)
	}
}

func CreateSubcategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			CategoryId  string `json:"category_id"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.CategoryId = strings.TrimSpace(body.CategoryId)

		name := body.Name
		description := body.Description
		categoryId := body.CategoryId

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedCategoryId, err := strconv.Atoi(categoryId)
		err = dbs.InsertSubcategory(name, description, parsedCategoryId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateSubcategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			CategoryId  string `json:"category_id"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Id = strings.TrimSpace(body.Id)
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.CategoryId = strings.TrimSpace(body.CategoryId)

		id := body.Id
		name := body.Name
		description := body.Description
		categoryId := body.CategoryId

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		parsedCategoryId, err := strconv.Atoi(categoryId)
		err = dbs.UpdateSubcategory(parsedId, name, description, parsedCategoryId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func DeleteSubcategory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		err = dbs.DeleteSubcategory(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetSubcategoriesByCategoryId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		subcategories, err := dbs.GetSubcategoriesByCategoryId(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(subcategories)
	}
}

func GetProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs := db.GetDbSource()
		products, err := dbs.GetProducts()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	}
}

func GetProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		product, err := dbs.GetProductById(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}

func GetProductsBySubcategoryId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		products, err := dbs.GetProductsBySubcategoryId(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	}
}

func CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Name             string `json:"name"`
			Description      string `json:"description"`
			SubcategoryId    string `json:"subcategory_id"`
			Price            string `json:"price"`
			CurrentInventory string `json:"current_inventory"`
			Image            string `json:"image"`
			Brand            string `json:"brand"`
			SKU              string `json:"sku"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.SubcategoryId = strings.TrimSpace(body.SubcategoryId)
		body.Price = strings.TrimSpace(body.Price)
		body.CurrentInventory = strings.TrimSpace(body.CurrentInventory)
		body.Image = strings.TrimSpace(body.Image)
		body.Brand = strings.TrimSpace(body.Brand)
		body.SKU = strings.TrimSpace(body.SKU)

		name := body.Name
		description := body.Description
		subcategoryId := body.SubcategoryId
		price := body.Price
		currentInventory := body.CurrentInventory
		image := body.Image
		brand := body.Brand
		sku := body.SKU

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedSubcategoryCategoryId, err := strconv.Atoi(subcategoryId)
		parsedPrice, err := strconv.ParseFloat(price, 64)
		parsedCurrentInventory, err := strconv.Atoi(currentInventory)
		err = dbs.InsertProduct(name, description, parsedSubcategoryCategoryId, parsedPrice, parsedCurrentInventory, image, brand, sku)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("failed to read body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Id               string `json:"id"`
			Name             string `json:"name"`
			Description      string `json:"description"`
			SubcategoryId    string `json:"subcategory_id"`
			Price            string `json:"price"`
			CurrentInventory string `json:"current_inventory"`
			Image            string `json:"image"`
			Brand            string `json:"brand"`
			SKU              string `json:"sku"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			fmt.Printf("failed to unmarshal body")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Id = strings.TrimSpace(body.Id)
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.SubcategoryId = strings.TrimSpace(body.SubcategoryId)
		body.Price = strings.TrimSpace(body.Price)
		body.CurrentInventory = strings.TrimSpace(body.CurrentInventory)
		body.Image = strings.TrimSpace(body.Image)
		body.Brand = strings.TrimSpace(body.Brand)
		body.SKU = strings.TrimSpace(body.SKU)

		id := body.Id
		name := body.Name
		description := body.Description
		subcategoryId := body.SubcategoryId
		price := body.Price
		currentInventory := body.CurrentInventory
		image := body.Image
		brand := body.Brand
		sku := body.SKU

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		parsedSubcategoryCategoryId, err := strconv.Atoi(subcategoryId)
		parsedPrice, err := strconv.ParseFloat(price, 64)
		parsedCurrentInventory, err := strconv.Atoi(currentInventory)
		err = dbs.UpdateProduct(parsedId, name, description, parsedSubcategoryCategoryId, parsedPrice, parsedCurrentInventory, image, brand, sku)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)
		err = dbs.DeleteProduct(parsedId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
