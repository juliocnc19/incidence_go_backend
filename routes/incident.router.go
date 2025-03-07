package routes

import (
	"incidence_grade/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpIncidentRouters(app *fiber.App, incidentHandler *handlers.IncidentHandler){
  incidents := app.Group("/incidents")
}
