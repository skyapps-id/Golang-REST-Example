package main

import (
	"golang-rest-example/controller"
	"golang-rest-example/model"
	"golang-rest-example/repository"
	"golang-rest-example/service"
	"golang-rest-example/shared"
	"golang-rest-example/shared/config"
	"golang-rest-example/shared/middlewares"
	"log"

	_ "golang-rest-example/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

// @title Book REST Example
// @version 1.0
// @description Documentation REST API Book
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /v1

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization
func main() {
	// Common Module
	deps := shared.NewDeps(cfg)
	middlewares := middlewares.NewMiddlewares()

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

	r.Any("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")

	book := v1.Group("/books", middlewares.Authenticate)
	{
		book.POST("", bookController.Create)
		book.GET("", bookController.Fatch)
		book.GET("/:id", bookController.FatchByID)
		book.PUT("/:id", bookController.Update)
		book.DELETE("/:id", bookController.Delete)
	}

	r.Run(":3000")
}
