package storage

import (
	"cth.release/go-rds/common"
	"cth.release/go-rds/rds"
	"github.com/gofiber/fiber/v2"
)

func SetChannelByMap(c *fiber.Ctx) error {
	var req SetChannelByMapDto
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(common.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	channel := c.Params("channel")
	key := c.Params("key")

	err := rds.GlobalRds.Set(channel, key, req.Data)

	if err != nil {
		return c.Status(500).JSON(common.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(common.BasicResponse{
		Success: true,
	})
}
