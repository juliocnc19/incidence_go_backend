package routes

import (
	dto "incidence_grade/dto/auth"
	"incidence_grade/use_case"
	"incidence_grade/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func SetUpAuthRouters(app *fiber.App, user *use_case.User) {
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
}
