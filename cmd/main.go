package main

import (
	"golang-rest-example/controller"
	"golang-rest-example/model"
	"golang-rest-example/repository"
	"golang-rest-example/service"
	"golang-rest-example/shared"
	"golang-rest-example/shared/config"
	"golang-rest-example/shared/middleware"
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	cfg *config.Configuration
)

func init() {
	var err error
	cfg, err = config.New()
	if err != nil {
		log.Println(err)
	}
	db, err = model.Database(cfg)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	// Common Module
	deps := shared.NewDeps(cfg)
	middleware := middleware.NewMiddlewares()

	// Book Module
	bookRepository := repository.NewBookRepository(db)
	bookService := service.NewBookService(bookRepository, deps)
	bookController := controller.NewBookController(bookService)

	r := gin.Default()

	cfgCors := cors.DefaultConfig()
	cfgCors.AllowAllOrigins = true
	cfgCors.AllowCredentials = true
	cfgCors.AddAllowHeaders("authorization")
	r.Use(cors.New(cfgCors))

	v1 := r.Group("/v1")

	book := v1.Group("/books", middleware.Authenticate)
	{
		book.POST("", bookController.Create)
		book.GET("", bookController.Fatch)
		book.GET("/:id", bookController.FatchByID)
		book.PUT("/:id", bookController.Update)
		book.DELETE("/:id", bookController.Delete)
	}

	r.Run(":3000")
}
