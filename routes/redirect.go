package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strings"
)

func Redirect(ctx *fiber.Ctx) error {

	code := ctx.Params("code")

	url, err := r.FindByCode(code)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).SendString("Not Found!")
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error!")
	}

	// TODO: UPDATE COUNTER WITH UPDATE STATEMENT +1

	redirect := url.Redirect
	if !strings.HasSuffix("http", strings.ToLower(redirect)) { // there should be proper handling and validation for URLs
		redirect = "http://" + redirect
	}
	return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
}
