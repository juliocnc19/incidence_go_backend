package routes

import (
	"incidence_grade/dto"
	"incidence_grade/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRouters(app *fiber.App, userHandler *handlers.UserHandler) {
	users := app.Group("/users")
	users.Get("/", func(c *fiber.Ctx) error {
		allUsers, error := userHandler.GetAllUsers()
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al obtener los usuarios",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   allUsers,
			"detail": "Usuarios obtenidos con exito",
			"length": len(allUsers),
		})
	})

	users.Post("/", func(c *fiber.Ctx) error {
		var input dto.CreateUserDto
		if error := c.BodyParser(&input); error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos invalidos",
				"detail": error.Error(),
			})
		}
		createdUser, err := userHandler.CreateUser(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear el usuario",
				"detail": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   createdUser,
			"detail": "Usuario Creado Correctamente",
		})
	})
}
