package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/jasminemutiara03/gocroot/controller"
)

func Web(page *fiber.App) {
	page.Get("/", controller.Homepage)                                    //ujicoba panggil package musik
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
}
