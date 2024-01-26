package main

import (
	"go_fiber/api"
	"go_fiber/bootstrap"
	"go_fiber/cmd/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/gofiber/swagger"
)

func main() {
	// Initialize the application's bootstrap components.
	app := bootstrap.NewInitializeBootstrap()

	// Create a new Fiber instance.
	f := fiber.New()

	f.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	f.Static("/static", "./")
	f.Get("/swagger/*", fiberSwagger.New(fiberSwagger.Config{
		URL: "http://localhost:3000/static/oapi_codegen.yml",
	}))

	// Serve the Swagger UI at the "/swagger" endpoint.
	// f.Static("/swagger", "cmd")

	// Define the API versioned group.
	f.Group("/api/v1.0")

	// Initialize the application-specific handlers.
	serve := handlers.NewServiceInitial(app)
	checkController := serve.CheckHandler()
	bookController := serve.BookHandler()

	// Create a wrapper for the controller interfaces.
	wrapper := &handlers.ServerInterfaceWrapper{
		CheckHandler: checkController,
		BookHandler:  bookController,
	}

	// Register API handlers.
	api.RegisterHandlers(f, wrapper)

	// Start the server on port 3000.
	f.Listen(":3000")
}
