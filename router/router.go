package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/ssr0016/book/controller"
)

func NewRouter(bookController *controller.BookController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("Hello World"))
	})

	router.POST("/book", bookController.Create)
	router.GET("/book/:bookID", bookController.GetByID)
	router.GET("/book", bookController.Search)
	router.PUT("/book/:bookID", bookController.Update)
	router.DELETE("/book/:bookID", bookController.Delete)

	return router
}
