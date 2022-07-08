package instance

import "github.com/gofiber/fiber/v2"

func AddRoutes(app *fiber.App) {
	ins := app.Group("/api/v1/instances")     // Instances
	ins.Get("/", getInstance)                 // /api/v1/instances GET
	ins.Get("/peers", getPeers)               // /api/v1/instances/peers GET
	ins.Get("/activity", getInstanceActivity) // /api/v1/instances/activity GET
}

func getInstance(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Instance",
	})
}

func getPeers(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Peers",
	})
}

func getInstanceActivity(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "InstanceActivity",
	})
}
