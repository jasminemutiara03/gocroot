package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jasminemutiara03/gocroot/config"

	"github.com/whatsauth/whatsauth"

	"github.com/gofiber/fiber/v2"
	"github.com/jasminemutiara03/gocroot/url"
)

func main() {
	go whatsauth.RunHub()
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
