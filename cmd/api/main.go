package main

import (
	"lentara-backend/internal/bootstrap"
	"log"
	"os"
	"time"
)

func main() {
	log.Printf("Current time: %v\n", time.Now())
	err := bootstrap.Start(os.Args)
	if err != nil {
		panic(err)
	}

	// _env, err := env.New()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// _mysql
	//
	// _fiber, err := fiber.New()
	// if err != nil {
	//   //
	// }
	//
	// cfg, err := env.New()
	// if err != nil {
	// 	panic(err)
	// }
	//
	//	app := fiber.New(fiber.Config{
	//		IdleTimeout: idleTimeout,
	//	})
	//
	// 	app.Get("/", func(c *fiber.Ctx) error {
	// 		return c.SendString("Hello world!")
	// 	})
	//
	// // Listen from a different goroutine
	//
	//	go func() {
	//		if err := app.Listen(":8080"); err != nil {
	//			log.Panic(err)
	//		}
	//	}()
	//
	// c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	// signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel
	//
	// _ = <-c // This blocks the main thread until an interrupt is received
	// fmt.Println("Gracefully shutting down...")
	// _ = app.Shutdown()
	//
	// fmt.Println("Running cleanup tasks...")
	//
	// fmt.Println("Fiber was successful shutdown.")
}
