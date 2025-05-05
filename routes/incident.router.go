package routes

import (
	"fmt"
	dto "incidence_grade/dto/incidents"
	"incidence_grade/middleware"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

const (
	UploadDirectory = "./uploads"
)

func SetUpIncidentRouters(app *fiber.App, incident *use_case.Incident) {
	incidents := app.Group("/incidents")

	incidents.Get("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {

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

	incidents.Get("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
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

	incidents.Post("/", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
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

	incidents.Put("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
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

	incidents.Delete("/:id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
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

	incidents.Get("/user/:user_id<int>", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		idUser := c.Params("user_id")
		idIncidentInt, _ := strconv.Atoi(idUser)
		usersIncidents, error := incident.FindByIdUser(uint(idIncidentInt))
		if error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error al obtener las incidencia",
				"detail": error.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   usersIncidents,
			"detail": "Incidencias de usario obtenidas",
			"length": len(usersIncidents),
		})
	})
	incidents.Post("/upload", middleware.JWTMiddleware, func(c *fiber.Ctx) error {
		incidentID := c.FormValue("incident_id")
		if incidentID == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Se requiere el incident_id",
				"detail": "Error al obtener el id de incidencia",
			})
		}

		idIncidentInt, _ := strconv.Atoi(incidentID)

		form, err := c.MultipartForm()
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error al procesar el formulario",
				"detail": err.Error(),
			})
		}

		files := form.File["files"]
		if len(files) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "No se encontraron archivos",
				"detail": "",
			})
		}

		var uploadedFiles []string
		for _, file := range files {
			filename := filepath.Join(UploadDirectory, file.Filename)
			if err := c.SaveFile(file, filename); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":  fmt.Sprintf("Error al guardar el archivo %s", file.Filename),
					"detail": err.Error(),
				})
			}
			uploadedFiles = append(uploadedFiles, filename)
		}

		fileCreated, error := incident.SaveFiles(uploadedFiles, uint(idIncidentInt))
		if error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Error registrar el archivo",
				"detail": error.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"data":   fileCreated,
			"detail": fmt.Sprintf("%d archivos subidos correctamente", len(files)),
		})
	})

	incidents.Get("/download/:filename", func(c *fiber.Ctx) error {
		filename := c.Params("filename")
		filePath := filepath.Join(UploadDirectory, filename)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error":  "Archivo no encontrado",
				"detail": err.Error(),
			})
		}
		return c.Download(filePath)
	})
}
