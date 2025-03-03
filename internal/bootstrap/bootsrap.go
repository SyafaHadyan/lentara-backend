package bootstrap

import (
	"fmt"
	producthandler "lentara-backend/internal/app/product/interface/rest"
	productrepository "lentara-backend/internal/app/product/repository"
	productusecase "lentara-backend/internal/app/product/usecase"
	productspecificationhandler "lentara-backend/internal/app/productspecification/interface/rest"
	productspecificationrepository "lentara-backend/internal/app/productspecification/repository"
	productspecificationusecase "lentara-backend/internal/app/productspecification/usecase"
	userhandler "lentara-backend/internal/app/user/interface/rest"
	userrepository "lentara-backend/internal/app/user/repository"
	userusecase "lentara-backend/internal/app/user/usecase"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/infra/fiber"
	"lentara-backend/internal/infra/jwt"
	"lentara-backend/internal/infra/mysql"
	"lentara-backend/internal/middleware"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Start(args []string) error {
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
	if err != nil {
		return err
	}

	if len(args) != 1 && args[1] == "--migrate" {
		err = mysql.Migrate(database)
		if err != nil {
			log.Println("failed to migrate")
			log.Println(err)
		} else {
			log.Println("migration success")
		}
	} else {
		log.Println("no migration performed, running normally ...")
	}

	log.Printf("args: %v\n", args)

	val := validator.New()

	app := fiber.New()

	jwt := jwt.NewJWT(config)

	middleware := middleware.NewMiddleWare(*jwt)

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		// AllowCredentials: true,
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	v1 := app.Group("/api/v1")

	productRepository := productrepository.NewProductMySQL(database)
	productUseCase := productusecase.NewProductUsecase(productRepository)
	producthandler.NewProductHandler(v1, val, productUseCase, middleware)
	productSpecificationRepository := productspecificationrepository.NewProductSpecificationMySQL(database)
	productSpecificationUseCase := productspecificationusecase.NewProductSpecificationUsecase(productSpecificationRepository)
	productspecificationhandler.NewProductSpecificationHandler(v1, val, productSpecificationUseCase)
	userRepository := userrepository.NewUserMySQL(database)
	userUseCase := userusecase.NewUserUsecase(userRepository, jwt)
	userhandler.NewUserHandler(v1, val, userUseCase)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
