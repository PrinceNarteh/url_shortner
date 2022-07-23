package handlers

import (
	"url_shortener/pkg/model"

	"github.com/gofiber/fiber/v2"
)

func GetAllRedirects(ctx *fiber.Ctx) error {
	urls, err := model.GetAllUrls()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all url links" + err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(urls)
}
