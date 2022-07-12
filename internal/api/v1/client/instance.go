package client

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type Instance struct {
	URI              string `json:"uri"`               // The domain name of the instance.
	Title            string `json:"title"`             // The title of the website.
	Description      string `json:"description"`       // Admin-defined description of the site.
	ShortDescription string `json:"short_description"` // A shorter description defined by the admin.
	Email            string `json:"email"`             // An email that may be contacted for any inquiries.
}

func addInstanceRoutes(app *fiber.App) {
	ins := app.Group("/api/v1/instance")      // Instances
	ins.Get("/", getInstance)                 // /api/v1/instances GET
	ins.Get("/peers", getPeers)               // /api/v1/instances/peers GET
	ins.Get("/activity", getInstanceActivity) // /api/v1/instances/activity GET
}

func getInstance(c *fiber.Ctx) error {
	instance := Instance{
		URI:              "https://example.com",
		Title:            "Example",
		Description:      "This is an example instance.",
		ShortDescription: "This is an example instance.",
		Email:            "abc@123.com",
	}

	if json, err := json.Marshal(instance); err != nil {
		return c.Status(500).JSON(map[string]string{
			"message": "Error marshalling instance",
		})
	} else {
		return c.Status(200).JSON(json)
	}
}

func getPeers(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}

func getInstanceActivity(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not implemented")
}
