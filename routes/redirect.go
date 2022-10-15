package routes

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
	"strconv"
	"strings"
	"url-shortner-service/repository"
	"url-shortner-service/utils"
)

// clicks increments configurations
var (
	maxIncrementClicksWorker = utils.GetEnvOrDefault("MAX_INCREMENT_WORKERS", "100")
	maxIncrementClicksQueue  = utils.GetEnvOrDefault("MAX_INCREMENT_QUEUE", "100000")
)
var maxQueueInt, _ = strconv.ParseUint(maxIncrementClicksQueue, 10, 32)
var maxWorkerInt, _ = strconv.ParseUint(maxIncrementClicksWorker, 10, 32)
var incrementClicksChn = make(chan uint64, maxQueueInt)

func startIncrementClicksWorker(incrementChan <-chan uint64) {
	for id := range incrementChan {
		r, err := repository.GetUrlRepositoryInstance()
		if err != nil {
			log.Printf("Couldn't get url repository to update clicks!, error: %s" + err.Error())
		}
		err = r.IncrementClicks(id)
		if err != nil {
			log.Printf("Error updating clicks" + err.Error())
		}
	}
}

var setup = false

func SetupIncrementWorkers() {
	if setup {
		return
	}
	setup = true

	for w := 0; w < int(maxWorkerInt); w++ {
		go startIncrementClicksWorker(incrementClicksChn)
	}
}

func Redirect(ctx *fiber.Ctx) error {

	code := ctx.Params("code")

	r, err := getRepositoryOrSendErr(ctx)
	if err != nil {
		return err
	}

	url, err := r.FindByCode(code)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).SendString("Not Found!")
		}
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error!")
	}

	// increment clicks in background
	incrementClicksChn <- url.ID

	redirect := url.Redirect
	if !strings.HasSuffix("http", strings.ToLower(redirect)) { // there should be proper handling and validation for URLs
		redirect = "http://" + redirect
	}
	return ctx.Redirect(redirect, fiber.StatusTemporaryRedirect)
}
