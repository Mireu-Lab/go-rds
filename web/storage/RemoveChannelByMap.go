package storage

import (
	"cth.release/go-rds/common"
	"cth.release/go-rds/rds"
	"github.com/gofiber/fiber/v2"
)

func RemoveChannelByMap(c *fiber.Ctx) error {
	channel := c.Params("channel")
	key := c.Params("key")

	err := rds.GlobalRds.Remove(channel, key)

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
