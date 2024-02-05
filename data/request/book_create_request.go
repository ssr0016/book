package request

type BookCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
