package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/book/data/request"
	"github.com/ssr0016/book/helper"
	"github.com/ssr0016/book/service"
)

type BookController struct {
	BookService service.BookService
}

func NewBookController(bookService service.BookService) *BookController {
	return &BookController{
		BookService: bookService,
	}
}

func (c *BookController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookCreateRequest := request.BookCreateRequest{}
	helper.ReadRequestBody(w, r, &bookCreateRequest)

	err := c.BookService.Create(r.Context(), bookCreateRequest)
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, bookCreateRequest)
}

func (c *BookController) GetByID(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookID := p.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	result := c.BookService.GetByID(r.Context(), int64(id))
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, result)
}

func (c *BookController) Search(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	result, err := c.BookService.Search(r.Context())
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, result)
}

func (c *BookController) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookUpdateRequest := request.BookUpdateRequest{}
	helper.ReadRequestBody(w, r, &bookUpdateRequest)

	bookID := p.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	bookUpdateRequest.ID = int64(id)

	err = c.BookService.Update(r.Context(), bookUpdateRequest)
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, bookUpdateRequest)
}

func (c *BookController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	bookID := p.ByName("bookID")
	id, err := strconv.Atoi(bookID)
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	err = c.BookService.Delete(r.Context(), int64(id))
	if err != nil {
		helper.WriteResponseBody(w, err.Error())
	}

	helper.WriteResponseBody(w, nil)
}
