package main

import (
	"fmt"
	"os"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create fiber app
	app := fiber.New()

	// Env
	app_id := os.Getenv("APP_ID")
	app_certificate := os.Getenv("APP_CERT")

	// Get RTC Token
	app.Get("/api/token/:channelName", func(c *fiber.Ctx) error {
		channelName := c.Params("channelName")
		uid := uint32(0)
		expireTimeInSeconds := uint32(3200)
		currentTimestamp := uint32(time.Now().UTC().Unix())
		expireTimestamp := expireTimeInSeconds + currentTimestamp

		result, err := rtctokenbuilder.BuildTokenWithUID(app_id, app_certificate, channelName, uid, rtctokenbuilder.RoleAttendee, expireTimestamp)
		if err != nil {
			fmt.Println(err)
			return c.Status(403).JSON(fiber.Map{"error": err})
		} else {
			return c.Status(200).JSON(fiber.Map{"token": result})
		}
	})

	// Run on port 8080
	app.Listen(":8080")
}
