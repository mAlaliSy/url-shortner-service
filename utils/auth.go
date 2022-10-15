package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetCurrentUserClaim(ctx *fiber.Ctx, claim string) interface{} {
	user := ctx.Locals("user").(*jwt.Token)
	return user.Claims.(jwt.MapClaims)[claim]
}

func GetCurrentUserId(ctx *fiber.Ctx) uint64 {
	return uint64(GetCurrentUserClaim(ctx, "id").(float64))
}
