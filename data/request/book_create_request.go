package request

import "github.com/ssr0016/book/data/response"

type BookCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

func (b BookCreateRequest) Validate() error {
	if len(b.Title) == 0 {
		return response.ErrNameEmpty
	}

	return nil
}
