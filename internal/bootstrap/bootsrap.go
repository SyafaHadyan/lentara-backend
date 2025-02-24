package bootstrap

import (
	"fmt"
	"lentara-backend/internal/infra/env"
	"lentara-backend/internal/infra/fiber"
	"lentara-backend/internal/infra/mysql"
)

func Start() error {
	config, err := env.New()
	if err != nil {
		panic(err)
	}

	_, err = mysql.New(fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPassword,
		config.DBName,
	))

	app := fiber.New()

	return app.Listen(fmt.Sprintf(":%d", config.AppPort))
}
