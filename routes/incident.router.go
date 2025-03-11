package routes

import (
	"incidence_grade/use_case"

	"github.com/gofiber/fiber/v2"
)

func SetUpIncidentRouters(app *fiber.App, incident *use_case.Incident){
  app.Group("/incidents")
}
