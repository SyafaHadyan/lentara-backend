package bootstrap

import (
	"fmt"
	producthandler "lentara-backend/internal/app/product/interface/rest"
	productrepository "lentara-backend/internal/app/product/repository"
	productusecase "lentara-backend/internal/app/product/usecase"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/infra/fiber"
	"lentara-backend/internal/infra/mysql"
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
		config.DBPassword,
		config.DBName,
	))

	app := fiber.New()

	v1 := app.Group("/api/v1")

	productRepository := productrepository.NewProductMySQL(database)
	productUseCase := productusecase.NewProductUsecase()
	producthandler.NewProductHandler(v1)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
