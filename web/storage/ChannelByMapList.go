package storage

import (
	"cth.release/go-rds/common"
	"cth.release/go-rds/rds"
	"github.com/gofiber/fiber/v2"
)

func ChannelByMapList(c *fiber.Ctx) error {
	channel := c.Params("channel")

	data, err := rds.GlobalRds.List(channel)

	if err != nil {
		return c.Status(500).JSON(common.BasicResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.Status(200).JSON(common.BasicResponse{
		Success: true,
		Data:    data,
	})
}
