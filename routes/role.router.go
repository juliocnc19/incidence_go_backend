package routes

import (
	dto "incidence_grade/dto/roles"
	"incidence_grade/middleware"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpRoleRouters(app *fiber.App, role *use_case.Role) {
	roles := app.Group("/roles")

	roles.Get("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		allRoles, err := role.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al obtener los roles",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   allRoles,
			"detail": "Roles obtenidos con éxito",
			"length": len(allRoles),
		})
	})

	roles.Get("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idRole := c.Params("id")
		idRoleInt, _ := strconv.Atoi(idRole)
		roleFind, err := role.GetById(uint(idRoleInt))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener el rol",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   roleFind,
			"detail": "Rol obtenido con éxito",
		})
	})

	roles.Post("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.CreateRoleDto

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

		createdRole, err := role.Create(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear el rol",
				"detail": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   createdRole,
			"detail": "Rol creado correctamente",
		})
	})

	roles.Put("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.UpdateRoleDto

		idRole := c.Params("id")
		idRoleInt, _ := strconv.Atoi(idRole)

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

		updatedRole, err := role.Update(uint(idRoleInt), input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al actualizar el rol",
				"detail": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   updatedRole,
			"detail": "Rol actualizado correctamente",
		})
	})

	roles.Delete("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		result, err := role.Delete(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar el rol",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   result,
			"detail": "Rol eliminado con éxito",
		})
	})
}
