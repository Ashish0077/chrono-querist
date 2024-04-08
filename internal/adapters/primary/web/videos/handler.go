package videos

import (
	"chrono-querist/internal/ports/primary/API"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Index(videosService API.VideosPort) fiber.Handler {
	return func(fiberContext *fiber.Ctx) error {
		log.Info("Index Action Called")
		pageNumber, err := strconv.Atoi(fiberContext.Query("page_number"))
		if err != nil {
			pageNumber = 1
		}
		limit, err := strconv.Atoi(fiberContext.Query("limit"))
		if err != nil {
			limit = 10
		}
		title := fiberContext.Query("title")
		description := fiberContext.Query("description")

		resp, err := videosService.Index(pageNumber, limit, &title, &description)
		if err != nil {
			return fiberContext.SendStatus(500)
		}

		return fiberContext.JSON(resp)
	}
}
