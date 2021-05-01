package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AgoraIO-Community/go-tokenbuilder/rtctokenbuilder"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("FATAL ERROR: Can't get ENV data")
	}
	appIDEnv := os.Getenv("APP_ID")
	appCertEnv := os.Getenv("APP_CERT")

	// Create fiber app
	app := fiber.New()

	// Get RTC Token
	app.Get("/api/token/:channelName", func(c *fiber.Ctx) error {
		channelName := c.Params("channelName")
		uid := uint32(0)
		fmt.Println(appIDEnv)
		fmt.Println(appCertEnv)
		fmt.Println(channelName)
		expireTimeInSeconds := uint32(3200)
		currentTimestamp := uint32(time.Now().UTC().Unix())
		expireTimestamp := expireTimeInSeconds + currentTimestamp

		result, err := rtctokenbuilder.BuildTokenWithUID(appIDEnv, appCertEnv, channelName, uid, rtctokenbuilder.RoleAttendee, expireTimestamp)
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
