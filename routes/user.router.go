package routes

import (
	"incidence_grade/dto/users"
	"incidence_grade/middleware"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpUserRouters(app *fiber.App, user *use_case.User) {
	users := app.Group("/users")

	// Get Users
	users.Get("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		allUsers, error := user.GetAll()
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

	//Get User for id:int
	users.Get("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idUser := c.Params("id")
		idUserInt, _ := strconv.Atoi(idUser)
		userFind, error := user.GetById(uint(idUserInt))
		if error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener usuario",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   userFind,
			"detail": "Usuario Obtenido",
		})
	})

	//Post User Create
	users.Post("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.CreateUserDto

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

		createdUser, error := user.Create(input)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear el usuario",
				"detail": error.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   createdUser,
			"detail": "Usuario Creado Correctamente",
		})
	})

	users.Put("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.UpdateUserDto

		idUser := c.Params("id")
		idUserInt, _ := strconv.Atoi(idUser)

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

		updatedUser, error := user.Update(uint(idUserInt), input)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al actualizar el usuario",
				"detail": error.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   updatedUser,
			"detail": "Usuario Actualizado Correctamente",
		})
	})

	// Cambiar contraseña
	users.Put("/:id<int>/change-password", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.ChangePasswordDto

		idUser := c.Params("id")
		idUserInt, _ := strconv.Atoi(idUser)

		errorParser := c.BodyParser(&input)
		if errorParser != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos inválidos",
				"detail": errorParser.Error(),
			})
		}

		error := utils.ValidateInput(input)
		if error != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validación",
				"detail": error.Error(),
			})
		}

		err := user.ChangePassword(uint(idUserInt), input.NewPassword)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al cambiar la contraseña",
				"detail": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"detail": "Contraseña actualizada correctamente",
		})
	})

	users.Delete("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		resutl, error := user.Delete(id)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar al usuario",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   resutl,
			"detail": "Usuario eliminado con exito",
		})
	})
}
