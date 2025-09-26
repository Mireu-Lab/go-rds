package channel

import (
	"cth.release/go-rds/common"
	"cth.release/go-rds/rds"
	"github.com/gofiber/fiber/v2"
)

func ChannelList(c *fiber.Ctx) error {
	return c.Status(200).JSON(common.BasicResponse{
		Success: true,
		Data:    rds.GlobalRds.ListChannels(),
	})
}
