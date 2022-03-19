package v1

import (
	"github.com/arcoz0308/arcoz0308.tech/routes/api/v1/minecraft"
	"github.com/gofiber/fiber/v2"
)

func Route(r fiber.Router) {
	// minecraft
	r.Get("/mc/:ip/full", minecraft.Full)
	r.Get("/mc/:ip", minecraft.Basic)
	r.Get("/mcbe/:ip", minecraft.Bedrock)
}
