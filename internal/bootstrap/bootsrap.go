package bootstrap

import (
	"fmt"
	producthandler "lentara-backend/internal/app/product/interface/rest"
	productrepository "lentara-backend/internal/app/product/repository"
	productusecase "lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/infra/fiber"
	"lentara-backend/internal/infra/mysql"

	"github.com/go-playground/validator/v10"
)

func Start() error {
	config, err := env.New()
	if err != nil {
		panic(err)
	}

	database, err := mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	))

	// TODO: Add flag to migrate (get from os args)
	err = mysql.Migrate(database)
	if err != nil {
		panic(err)
	}

	val := validator.New()

	app := fiber.New()

	v1 := app.Group("/api/v1")

	productRepository := productrepository.NewProductMySQL(database)
	productUseCase := productusecase.NewProductUsecase(productRepository)
	producthandler.NewProductHandler(v1, val, productUseCase)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
