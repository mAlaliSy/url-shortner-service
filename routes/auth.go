package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
	"time"
	"url-shortner-service/entity"
	"url-shortner-service/repository"
	"url-shortner-service/utils"
)

func getUserRepositoryOrSendErr(ctx *fiber.Ctx) (*repository.UserRepositoryImpl, error) {
	var r, err = repository.GetUserRepositoryInstance()
	if err != nil {
		log.Printf("Couldn't get User repository, error: %s", err.Error())
		_ = ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return r, err
}

var jwtSecret = utils.GetEnvOrDefault("JWT_SECRET", "ThIsIsFoRtEsTiNgOnLy")

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(ctx *fiber.Ctx) error {

	var credentials AuthRequest

	err := ctx.BodyParser(&credentials)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials"})
	}

	r, err := getUserRepositoryOrSendErr(ctx)
	if err != nil {
		return err
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
		log.Printf("Couldn't sign jwt token, error: %s", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong!"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": stoken,
	})
}

func Register(ctx *fiber.Ctx) error {
	var request AuthRequest
	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request!"})
	}

	if len(request.Username) < 3 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Username length must be at least 3 characters!"})
	}

	if len(request.Password) < 6 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Password length must be at least 6 characters!"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Couldn't hash password, error: %s", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong!"})
	}

	r, err := getUserRepositoryOrSendErr(ctx)
	if err != nil {
		return err
	}
	var user = entity.User{Username: request.Username, Password: string(hashedPassword)}
	err = r.Create(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Username already used!",
			})
		}
		log.Printf("Couldn't register user, error: %s", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Something went wrong!"})
	}

	return ctx.Status(fiber.StatusOK).JSON(&user)
}
