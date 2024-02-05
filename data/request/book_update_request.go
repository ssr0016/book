package request

type BookUpdateRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
