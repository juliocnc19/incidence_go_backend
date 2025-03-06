package main

import (
	"incidence_grade/config"
	"incidence_grade/handlers"
	"incidence_grade/repository"
	"incidence_grade/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
  app.Use(logger.New(logger.Config{
    Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
  }))

	environments := config.LoadEnviroments()
	db := config.InitDB(environments)
	userRepo := repository.NewUserRepository(db)
	userHanlder := handlers.NewUserHandler(userRepo)
	routes.SetUpUserRouters(app, userHanlder)

	app.Listen(":3001")
}
