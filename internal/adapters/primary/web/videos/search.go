package videos

import (
	"chrono-querist/internal/core/domain/video"
	"chrono-querist/internal/core/services/videos"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Search(videosAPI videos.API) fiber.Handler {
	return func(fctx *fiber.Ctx) error {
		ctx := fctx.Context()

		fmt.Println("DOING SOMETHING")
		query := fctx.Query("query")

		fmt.Printf("%v", query)

		resp, err := videosAPI.Search(ctx, query)
		if err != nil {
			return fctx.SendStatus(500)
		}

		var out struct {
			Videos []video.Video `json:"videos"`
		}
		out.Videos = resp
		return fctx.JSON(out)
	}
}
