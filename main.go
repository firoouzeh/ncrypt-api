package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"ncrypt-api/config"
	"ncrypt-api/handlers"
)

func main() {
	c := config.BuildConfig()

	di, err := handlers.BuildDI(c)
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit("100K"))
	//e.Use(middleware.Logger())

	customValidator := validator.New()
	e.Validator = &Validator{Validator: customValidator}

	e.GET("/", di.Index)
	e.GET("/api/v1", di.GetIndexV1)
	e.POST("/api/v1/note", di.PostStoreSecureNoteV1)
	e.GET("/api/v1/note/:id", di.GetSecureNoteV1)

	e.Logger.Fatal(e.Start(":1990"))
}

//Validate echo custom validator
func (v Validator) Validate(i interface{}) error {
	return v.Validator.Struct(i)
}
