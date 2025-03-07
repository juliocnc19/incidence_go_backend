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
    Format: "[${ip}] ${status} - ${method} ${path} - ${latency}\n",
  }))
  
  //db
	environments := config.LoadEnviroments()
	db := config.InitDB(environments)
  
  //Repository
	userRepo := repository.NewUserRepository(db)
  incidentRepo := repository.NewIncidentRepository(db)
  
  //Handler
	userHandler := handlers.NewUserHandler(userRepo)
  handlers.NewIncidentHandler(incidentRepo)
  
  //Routers
	routes.SetUpUserRouters(app, userHandler)

	app.Listen(":3001")
}
