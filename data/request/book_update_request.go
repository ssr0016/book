package request

import "github.com/ssr0016/book/data/response"

type BookUpdateRequest struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (b BookUpdateRequest) Validate() error {
	if len(b.Title) == 0 {
		return response.ErrNameEmpty
	}

	return nil
}
