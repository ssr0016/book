package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ssr0016/book/config"
	"github.com/ssr0016/book/controller"
	"github.com/ssr0016/book/repository"
	"github.com/ssr0016/book/router"
	"github.com/ssr0016/book/service"
)

func main() {
	fmt.Printf("Start server")

	// database
	db := config.DatabaseConnection()

	// repository
	bookRepository := repository.NewBookRepository(db)

	// service
	bookService := service.NewBookRepositoryImpl(bookRepository)

	// controller
	bookController := controller.NewBookController(bookService)

	// router
	routes := router.NewRouter(bookController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
