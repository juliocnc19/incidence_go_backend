package routes

import (
	dto "incidence_grade/dto/attachments"
	"incidence_grade/middleware"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SetUpAttachmentRouters(app *fiber.App, attachment *use_case.Attachment) {
	attachments := app.Group("/attachments")

	attachments.Get("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		allAttachments, err := attachment.GetAll()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al obtener los archivos adjuntos",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   allAttachments,
			"detail": "Archivos adjuntos obtenidos con éxito",
			"length": len(allAttachments),
		})
	})

	attachments.Get("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idAttachment := c.Params("id")
		idAttachmentInt, _ := strconv.Atoi(idAttachment)
		attachmentFind, err := attachment.GetById(uint(idAttachmentInt))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener el archivo adjunto",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   attachmentFind,
			"detail": "Archivo adjunto obtenido con éxito",
		})
	})

	attachments.Get("/incident/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idIncident := c.Params("id")
		idIncidentInt, _ := strconv.Atoi(idIncident)
		attachmentsFind, err := attachment.GetByIncidentId(uint(idIncidentInt))
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener los archivos adjuntos del incidente",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   attachmentsFind,
			"detail": "Archivos adjuntos del incidente obtenidos con éxito",
		})
	})

	attachments.Post("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.CreateAttachmentDto

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

		createdAttachment, err := attachment.Create(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al crear el archivo adjunto",
				"detail": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   createdAttachment,
			"detail": "Archivo adjunto creado correctamente",
		})
	})

	attachments.Put("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		var input dto.UpdateAttachmentDto

		idAttachment := c.Params("id")
		idAttachmentInt, _ := strconv.Atoi(idAttachment)

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

		updatedAttachment, err := attachment.Update(uint(idAttachmentInt), input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al actualizar el archivo adjunto",
				"detail": err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   updatedAttachment,
			"detail": "Archivo adjunto actualizado correctamente",
		})
	})

	attachments.Delete("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		id, _ := strconv.Atoi(c.Params("id"))
		result, err := attachment.Delete(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar el archivo adjunto",
				"detail": err.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   result,
			"detail": "Archivo adjunto eliminado con éxito",
		})
	})
} 
