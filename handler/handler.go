package handler

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"vayer-electric-backend/db"
	"vayer-electric-backend/logging"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

var (
	volumePath = "./uploads/"
	log        = logging.GetLogger()
)

func GetCategories() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbs := db.GetDbSource()
		categories, err := dbs.GetCategories()

		if err != nil {
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		dbs := db.GetDbSource()
		category, err := dbs.GetCategoryById(parsedId)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
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
			log.Error(err.Error())
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
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			ImageUrl    string `json:"image_url"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.ImageUrl = strings.TrimSpace(body.ImageUrl)

		name := body.Name
		description := body.Description
		image_url := body.ImageUrl

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dbs := db.GetDbSource()
		err = dbs.InsertCategory(name, description, image_url)

		if err != nil {
			log.Error(err.Error())
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
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			ImageUrl    string `json:"image_url"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Id = strings.TrimSpace(body.Id)
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.ImageUrl = strings.TrimSpace(body.ImageUrl)

		id := body.Id
		name := body.Name
		description := body.Description
		image_url := body.ImageUrl

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.UpdateCategory(parsedId, name, description, image_url)

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

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.DeleteCategory(parsedId)

		if err != nil {
			log.Error(err.Error())
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
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		subcategory, err := dbs.GetSubcategoryById(parsedId)

		if err != nil {
			log.Error(err.Error())
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
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Name        string `json:"name"`
			Description string `json:"description"`
			CategoryId  string `json:"category_id"`
			ImageUrl    string `json:"image_url"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.CategoryId = strings.TrimSpace(body.CategoryId)
		body.ImageUrl = strings.TrimSpace(body.ImageUrl)

		name := body.Name
		description := body.Description
		categoryId := body.CategoryId
		image_url := body.ImageUrl

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedCategoryId, err := strconv.Atoi(categoryId)

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.InsertSubcategory(name, description, parsedCategoryId, image_url)

		if err != nil {
			log.Error(err.Error())
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
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var body struct {
			Id          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			CategoryId  string `json:"category_id"`
			ImageUrl    string `json:"image_url"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Id = strings.TrimSpace(body.Id)
		body.Name = strings.TrimSpace(body.Name)
		body.Description = strings.TrimSpace(body.Description)
		body.CategoryId = strings.TrimSpace(body.CategoryId)
		body.ImageUrl = strings.TrimSpace(body.ImageUrl)

		id := body.Id
		name := body.Name
		description := body.Description
		categoryId := body.CategoryId
		image_url := body.ImageUrl

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)

		if err != nil {
			log.Error(err.Error())
			return
		}

		parsedCategoryId, err := strconv.Atoi(categoryId)

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.UpdateSubcategory(parsedId, name, description, parsedCategoryId, image_url)

		if err != nil {
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.DeleteSubcategory(parsedId)

		if err != nil {
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		subcategories, err := dbs.GetSubcategoriesByCategoryId(parsedId)

		if err != nil {
			log.Error(err.Error())
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
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		product, err := dbs.GetProductById(parsedId)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(product)
	}
}

func GetProductByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		dbs := db.GetDbSource()
		product, err := dbs.GetProductByName(name)

		if err != nil {
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		products, err := dbs.GetProductsBySubcategoryId(parsedId)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	}
}

func GetProductsByCategoryId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		products, err := dbs.GetProductsByCategoryId(parsedId)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	}
}

func GetProductsByCategoryName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		dbs := db.GetDbSource()
		products, err := dbs.GetProductsByCategoryName(name)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(products)
	}
}

func ServeProductImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")

		imagePath := volumePath + name

		http.ServeFile(w, r, imagePath)
	}
}

func CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20) // Limit to 10 MB file size
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Extract fields from the form
		name := r.FormValue("name")
		description := r.FormValue("description")
		subcategory := r.FormValue("subcategory")
		price := r.FormValue("price")
		currentInventory := r.FormValue("current_inventory")
		brand := r.FormValue("brand")
		sku := r.FormValue("sku")

		// Process the image file
		imageFile, _, err := r.FormFile("image")
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer imageFile.Close()

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()

		subcategoryObj, err := dbs.GetSubcategoryByName(subcategory)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		parsedPrice, err := strconv.ParseFloat(price, 64)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		parsedCurrentInventory, err := strconv.Atoi(currentInventory)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate a unique name for the image using the timestamp
		imageName, err := generateRandomFilename(".jpg", 10)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		imagePath := volumePath + imageName
		newImageFile, err := os.Create(imagePath)
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer newImageFile.Close()

		_, err = io.Copy(newImageFile, imageFile) // Copy image data
		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dbs = db.GetDbSource()

		err = dbs.InsertProduct(name, description, int(subcategoryObj.Id), parsedPrice, parsedCurrentInventory, imageName, brand, sku)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func generateRandomFilename(extension string, length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	for i := 0; i < length; i++ {
		randomBytes[i] = charset[randomBytes[i]%byte(len(charset))]
	}

	return string(randomBytes) + extension, nil
}

func UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw, err := io.ReadAll(r.Body)
		if err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		id := chi.URLParam(r, "id")

		var body struct {
			Name             string  `json:"name"`
			Price            float64 `json:"price"`
			CurrentInventory int     `json:"current_inventory"`
		}
		if err := json.Unmarshal(raw, &body); err != nil {
			log.Error(err.Error())
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Trim input
		body.Name = strings.TrimSpace(body.Name)

		name := body.Name
		price := body.Price
		currentInventory := body.CurrentInventory

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		dbs := db.GetDbSource()
		parsedId, err := strconv.Atoi(id)

		if err != nil {
			log.Error(err.Error())
			return
		}

		log.Info("Updating product with id: ", zap.Int("id", parsedId), zap.String("name", name), zap.Float64("price", price), zap.Int("current_inventory", currentInventory))

		err = dbs.UpdateProduct(parsedId, name, price, currentInventory)

		if err != nil {
			log.Error(err.Error())
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

		if err != nil {
			log.Error(err.Error())
			return
		}

		err = dbs.DeleteProduct(parsedId)

		if err != nil {
			log.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
