package main

import (
	"incidence_grade/config"
	"incidence_grade/repository"
	"incidence_grade/routes"
	"incidence_grade/use_case"

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
	incident := use_case.NewIncident(incidentRepo)

	//Routers
	routes.SetUpUserRouters(app, user)
	routes.SetUpIncidentRouters(app, incident)

	app.Listen(":3001")
}
