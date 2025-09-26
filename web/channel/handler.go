package channel

import "github.com/gofiber/fiber/v2"

func channelMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(api fiber.Router) {
	channel := api.Group("/channel", channelMiddleware)

	channel.Get("/list", ChannelList)
}
