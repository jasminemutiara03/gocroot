package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jasminemutiara03/gocroot/controller"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest) //API from user whatsapp message from iteung gowa

}
