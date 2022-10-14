package routes

import (
	"github.com/gofiber/fiber/v2"
	"url-shortner-service/repository"
)

func GetAll(ctx *fiber.Ctx) error {
	r := repository.GetUrlRepositoryInstance()
	all, err := r.GetAll()
	if err != nil {
		//log.Print("ERROR::", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(all)
}
