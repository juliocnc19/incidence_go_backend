package routes

import (
	dto "incidence_grade/dto/auth"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func SetUpAuthRouters(app *fiber.App, user *use_case.User, userToken *use_case.UserToken) {
	auth := app.Group("/auth")

	auth.Post("/", func(c *fiber.Ctx) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"iat": time.Now().Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "no se pudo generar el token",
			})
		}
		var input dto.LoginUserDto

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
		login, error := user.Login(input)

		if error != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error":  "Credenciales invalidas",
				"detail": error.Error(),
			})
		}
		return c.JSON(fiber.Map{
			"data":   login,
			"detail": "Autenticacion exitosa",
			"token":  tokenString,
		})
	})

	auth.Post("/register", func(c *fiber.Ctx) error {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp": time.Now().Add(time.Hour * 24).Unix(),
			"iat": time.Now().Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "no se pudo generar el token",
			})
		}
		var input dto.RegisterUserDto

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

		register, error := user.Register(input)
		if error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al registrar el usuario",
				"detail": error.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   register,
			"detail": "Usuario registrado",
			"token":  tokenString,
		})
	})

	auth.Post("/device-token", func(c *fiber.Ctx) error {
		var input dto.DeviceTokenDto

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Datos inválidos",
				"detail": err.Error(),
			})
		}

		if err := utils.ValidateInput(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":  "Error de validación",
				"detail": err.Error(),
			})
		}

		token, err := userToken.SaveDeviceToken(input)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al guardar el token del dispositivo",
				"detail": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"data":   token,
			"detail": "Token guardado",
		})
	})

	auth.Delete("/device-token/:id", func(c *fiber.Ctx) error {
		id, err := strconv.ParseUint(c.Params("id"), 10, 32)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "ID de token inválido",
			})
		}

		err = userToken.DeleteDeviceToken(uint(id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":  "Error al eliminar el token del dispositivo",
				"detail": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data":   nil,
			"detail": "Token de dispositivo eliminado correctamente",
		})
	})
}
