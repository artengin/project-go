package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

const profileUnknown = "unknown"

func main() {
	webApp := fiber.New()
	// Обозначаем, что на GET запрос по пути /address нужно вернуть строку с адресом
	webApp.Get("/address", func(c *fiber.Ctx) error {
		return c.SendString("145 DUNDEE SOUTH SAN FRANCISCO CA 94080-1023. USA")
	})

	webApp.Get("/profiles", func(c *fiber.Ctx) error {
		profileID := c.Query("profile_id", profileUnknown)
		if profileID == "" {
			profileID = profileUnknown
		}

		if profileID == profileUnknown {
			return c.Status(fiber.StatusUnprocessableEntity).SendString("profile_id is required")
		}

		return c.SendString(fmt.Sprintf("User Profile ID: %s", profileID))
	})


    webApp.Get("/counter/:event", func(c *fiber.Ctx) error {
        event := c.Params("event", "")
        if event == "" {
            return c.SendStatus(fiber.StatusUnprocessableEntity)
        }
        return c.SendString(event)
    })
	// Запускаем веб-приложение на порту 80
	// Оборачиваем в функцию логирования, чтобы видеть ошибки, если они возникнут
	logrus.Fatal(webApp.Listen(":8080"))
}
