package routes

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
	"strings"
	"url-shortner-service/entity"
	"url-shortner-service/repository"
	"url-shortner-service/utils"
)

var r = repository.GetUrlRepositoryInstance()

func GetAll(ctx *fiber.Ctx) error {
	all, err := r.GetAll()
	if err != nil {
		//log.Print("ERROR::", err.Error())
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&all)
}

func Get(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id",
		})
	}
	url, err := r.Get(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&url)
}

func Create(ctx *fiber.Ctx) error {
	var url entity.Url
	err := ctx.BodyParser(&url)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request!",
		})
	}
	generate := url.Code == ""
regenerate:
	if generate {
		url.Code = utils.RandomAlphanumeric(6)
	}
	url.Clicks = 0
	err = r.Create(&url)
	if err != nil {
		if generate {
			// generated code already used, regenerate a new one -- not an elegant way
			goto regenerate
		}
		if strings.Contains(err.Error(), "duplicate key") {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Code already used!",
			})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong!",
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(&url)
}

func Delete(ctx *fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid id!",
		})
	}
	err = r.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Something went wrong!",
		})
	}
	ctx.Status(fiber.StatusNoContent)
	return nil
}
