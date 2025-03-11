package main

import (
	"incidence_grade/config"
	"incidence_grade/use_case"
	"incidence_grade/repository"
	"incidence_grade/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
  app.Use(logger.New(logger.Config{
    Format: "[${ip}] ${status} ${method} ${path} ${latency}\n",
  }))
  
  //db
	environments := config.LoadEnviroments()
	db := config.InitDB(environments)
  
  //Repository
	userRepo := repository.NewUserRepository(db)
  incidentRepo := repository.NewIncidentRepository(db)
  
  //Handler
	user := use_case.NewUser(userRepo)
  use_case.NewIncident(incidentRepo)
  
  //Routers
	routes.SetUpUserRouters(app, user)

	app.Listen(":3001")
}
