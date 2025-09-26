package storage

import "github.com/gofiber/fiber/v2"

func storageMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(api fiber.Router) {
	storage := api.Group("/storage", storageMiddleware)

	storage.Get("/:channel/data/:key", ChannelByMap)

	storage.Get("/:channel/list", ChannelByMapList)

	storage.Post("/:channel/:key/set", SetChannelByMap)

	storage.Delete("/:channel/:key/remove", RemoveChannelByMap)
}
