package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"url-shortner-service/repository"
	"url-shortner-service/utils"
)

var jwtSecret = utils.GetEnvOrDefault("JWT_SECRET", "ThIsIsFoRtEsTiNgOnLy")

func Login(ctx *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var credentials Request

	err := ctx.BodyParser(&credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	r, err := repository.GetUserRepositoryInstance()
	if err != nil {
		log.Printf("Couldn't get user repository, error: %s", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong!"})
	}
	user, err := r.FindByUsername(credentials.Username)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	// Generate JWT
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.Username
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(15 * time.Minute)
	stoken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong!"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": stoken,
	})
}
