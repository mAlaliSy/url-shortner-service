package middleware

import (
	"github.com/gofiber/fiber"
	fiber2 "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"log"
	"url-shortner-service/utils"
)

var jwtSecret = utils.GetEnvOrDefault("JWT_SECRET", "ThIsIsFoRtEsTiNgOnLy")

func AuthRequired() fiber2.Handler {
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(ctx *fiber2.Ctx, err error) error {
			log.Printf("Authentication error: %s", err.Error())
			return ctx.SendStatus(fiber.StatusUnauthorized)
		},
		SigningKey: []byte(jwtSecret),
	})
}
