package response

type BookResponse struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}
