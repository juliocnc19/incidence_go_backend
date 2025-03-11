package routes

import (
	dto "incidence_grade/dto/incidents"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpIncidentRouters(app *fiber.App, incident *use_case.Incident) {
	incidents := app.Group("/incidents")

	incidents.Get("/", func(c *fiber.Ctx) error {

		allIncidents, error := incident.GetAll()
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al obtener las incidencias",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   allIncidents,
			"detail": "Incidencias obtenidos con exito",
			"length": len(allIncidents),
		})
	})

	incidents.Get("/:id<int>", func(c *fiber.Ctx) error {
		idIncident := c.Params("id")
		idIncidentInt, _ := strconv.Atoi(idIncident)
		incidentFind, error := incident.GetById(uint(idIncidentInt))
		if error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener la incidencia",
				"detail": error.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   incidentFind,
			"detail": "Incidencia obtenida",
		})
	})

	incidents.Post("/", func(c *fiber.Ctx) error {
		var input dto.CreateIncidentDto
		errorParser := c.BodyParser(&input)
		if errorParser != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos invalidos",
				"detail": errorParser.Error(),
			})
		}

		error := utils.ValidateInput(input)
		if error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validacion",
				"detail": error.Error(),
			})
		}

		incidentCreated, error := incident.Create(input)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear la incidencia",
				"detail": error.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   incidentCreated,
			"detail": "Incidencia creada",
		})

	})

	incidents.Put("/:id<int>", func(c *fiber.Ctx) error {
		var input dto.UpdateIncidentDto
		idIncident := c.Params("id")
		idIncidentInt, _ := strconv.Atoi(idIncident)

		errorParser := c.BodyParser(&input)
		if errorParser != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos invalidos",
				"detail": errorParser.Error(),
			})
		}

		error := utils.ValidateInput(input)
		if error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validacion",
				"detail": error.Error(),
			})
		}

		incidentUpdated, error := incident.Update(uint(idIncidentInt), input)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al actualizar la incidencia",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   incidentUpdated,
			"detail": "Incidencia actualizada",
		})
	})

	incidents.Delete("/:id<int>", func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		resutl, error := incident.Delete(uint(id))
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar la incidencia",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   resutl,
			"detail": "Incidencia eliminada con exito",
		})
	})
}
