package structs

type Category struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	ImageUrl    string `json:"image_url"`
}
