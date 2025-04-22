package routes

import (
	dto "incidence_grade/dto/statuses"
	"incidence_grade/middleware"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpStatusRouters(app *fiber.App, status *use_case.Status) {
	statuses := app.Group("/statuses")

	statuses.Get("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		allStatuses, err := status.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al obtener los estados",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   allStatuses,
			"detail": "Estados obtenidos con éxito",
			"length": len(allStatuses),
		})
	})

	statuses.Get("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idStatus := c.Params("id")
		idStatusInt, _ := strconv.Atoi(idStatus)
		statusFind, err := status.GetById(uint(idStatusInt))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener el estado",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   statusFind,
			"detail": "Estado obtenido con éxito",
		})
	})

	statuses.Post("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.CreateStatusDto

		err := c.BodyParser(&input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos inválidos",
				"detail": err.Error(),
			})
		}

		err = utils.ValidateInput(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validación",
				"detail": err.Error(),
			})
		}

		createdStatus, err := status.Create(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear el estado",
				"detail": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   createdStatus,
			"detail": "Estado creado correctamente",
		})
	})

	statuses.Put("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.UpdateStatusDto

		idStatus := c.Params("id")
		idStatusInt, _ := strconv.Atoi(idStatus)

		err := c.BodyParser(&input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos inválidos",
				"detail": err.Error(),
			})
		}

		err = utils.ValidateInput(input)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validación",
				"detail": err.Error(),
			})
		}

		updatedStatus, err := status.Update(uint(idStatusInt), input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al actualizar el estado",
				"detail": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   updatedStatus,
			"detail": "Estado actualizado correctamente",
		})
	})

	statuses.Delete("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		result, err := status.Delete(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar el estado",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   result,
			"detail": "Estado eliminado con éxito",
		})
	})
}

