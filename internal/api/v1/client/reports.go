package client

import "github.com/gofiber/fiber/v2"

func GetReports(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Reports",
	})
}

func GetReport(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "Report",
	})
}

func AssignReportToSelf(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "AssignReportToSelf",
	})
}

func UnassignReport(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "UnassignReport",
	})
}

func ResolveReport(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "ResolveReport",
	})
}

func ReopenReport(c *fiber.Ctx) error {
	return c.Status(200).JSON(map[string]string{
		"message": "ReopenReport",
	})
}
