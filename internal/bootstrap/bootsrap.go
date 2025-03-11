package bootstrap

import (
	"fmt"
	carthandler "lentara-backend/internal/app/cart/interface/rest"
	cartrepository "lentara-backend/internal/app/cart/repository"
	cartusecase "lentara-backend/internal/app/cart/usecase"
	producthandler "lentara-backend/internal/app/product/interface/rest"
	productrepository "lentara-backend/internal/app/product/repository"
	productusecase "lentara-backend/internal/app/product/usecase"
	productmediahandler "lentara-backend/internal/app/productmedia/interface/rest"
	productmediarepository "lentara-backend/internal/app/productmedia/repository"
	productmediausecase "lentara-backend/internal/app/productmedia/usecase"
	productspecificationhandler "lentara-backend/internal/app/productspecification/interface/rest"
	productspecificationrepository "lentara-backend/internal/app/productspecification/repository"
	productspecificationusecase "lentara-backend/internal/app/productspecification/usecase"
	sellerhandler "lentara-backend/internal/app/seller/interface/rest"
	sellerrepository "lentara-backend/internal/app/seller/repository"
	sellerusecase "lentara-backend/internal/app/seller/usecase"
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
	"github.com/gofiber/fiber/v2/middleware/idempotency"
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

	app.Use(
		cors.New(cors.Config{
			// AllowHeaders: "Authorization,Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
			AllowHeaders: "*",
			AllowOrigins: "*",
			// AllowCredentials: true,
			AllowMethods: "*",
		}),
		idempotency.New(),
	)

	v1 := app.Group("/api/v1")

	productSpecificationRepository := productspecificationrepository.NewProductSpecificationMySQL(database)
	productSpecificationUseCase := productspecificationusecase.NewProductSpecificationUsecase(productSpecificationRepository)
	productspecificationhandler.NewProductSpecificationHandler(v1, val, productSpecificationUseCase)
	userRepository := userrepository.NewUserMySQL(database)
	userUseCase := userusecase.NewUserUseCase(userRepository, jwt)
	userhandler.NewUserHandler(v1, val, middleware, userUseCase)
	productMediaRepository := productmediarepository.NewProductMediaMySQL(database)
	productMediaUseCase := productmediausecase.NewProductMediaUsecase(productMediaRepository)
	productmediahandler.NewProductMediahandler(v1, productMediaUseCase)
	sellerRepository := sellerrepository.NewSellerMySQL(database)
	sellerUseCase := sellerusecase.NewSellerUsecase(sellerRepository, jwt)
	sellerhandler.NewSellerHandler(v1, val, sellerUseCase)
	productRepository := productrepository.NewProductMySQL(database)
	productUseCase := productusecase.NewProductUseCase(productRepository)
	producthandler.NewProductHandler(v1, val, middleware, productUseCase, sellerUseCase)
	cartRepository := cartrepository.NewCartMySQL(database)
	cartUseCase := cartusecase.NewCartUsecase(cartRepository)
	carthandler.NewCartHandler(v1, val, middleware, cartUseCase, userUseCase, productUseCase)

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
