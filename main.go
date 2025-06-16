package main

import (
	"incidence_grade/config"
	"incidence_grade/repository"
	"incidence_grade/routes"
	"incidence_grade/use_case"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	if err := config.CreateDirectory(); err != nil {
		log.Fatalf("Error inicializando directorio de uploads: %v", err)
	}
	app := fiber.New(fiber.Config{
		BodyLimit: 100 * 1024 * 1024,
	})
	app.Use(logger.New(logger.Config{
		Format: "[${ip}] ${status} ${method} ${path} ${latency}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	//db
	environments := config.LoadEnviroments()
	db := config.InitDB(environments)

	//Repository
	userRepo := repository.NewUserRepository(db)
	incidentRepo := repository.NewIncidentRepository(db, environments)
	roleRepo := repository.NewRoleRepository(db)
	statusRepo := repository.NewStatusRepository(db)
	attachmentRepo := repository.NewAttachmentRepository(db)
	userTokenRepo := repository.NewUserTokenRepository(db)

	//Handler
	user := use_case.NewUser(userRepo)
	incident := use_case.NewIncident(incidentRepo)
	role := use_case.NewRole(roleRepo)
	status := use_case.NewStatus(statusRepo)
	attachment := use_case.NewAttachment(attachmentRepo)
	userToken := use_case.NewUserToken(userTokenRepo)

	//Routers
	routes.SetUpUserRouters(app, user)
	routes.SetUpIncidentRouters(app, incident)
	routes.SetUpAuthRouters(app, user, userToken)
	routes.SetUpRoleRouters(app, role)
	routes.SetUpStatusRouters(app, status)
	routes.SetUpAttachmentRouters(app, attachment)

	app.Listen("0.0.0.0:3001")
}
