package middleware

import (
	"fiber-demo-sftp/model"
	"fiber-demo-sftp/util"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadluth/log"
	"github.com/panjf2000/ants/v2"
)

type Middleware struct {
	properties     model.Properties
	poolConnection *ants.Pool
}

func NewMiddleware(properties model.Properties) Middleware {
	poolConnection, _ := ants.NewPool(int(properties.PoolSize), ants.WithPreAlloc(true))
	return Middleware{properties, poolConnection}
}

func (m *Middleware) ServiceMiddleware() func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		var traceId = util.CreateUniqID()
		m.poolConnection.Submit(func() {
			log.Message(
				traceId,
				"IN",
				"GO-FIBER",
				"",
				"URL",
				ctx.OriginalURL(),
				"",
				"REQUEST",
				string(ctx.Request().Body()))
		})

		ctx.Set("X-XSS-Protection", "1; mode=block")
		ctx.Set("X-Content-Type-Options", "nosniff")
		ctx.Set("X-Download-Options", "noopen")
		ctx.Set("Strict-Transport-Security", "max-age=5184000")
		ctx.Set("X-Frame-Options", "SAMEORIGIN")
		ctx.Set("X-DNS-Prefetch-Control", "off")

		ctx.Locals("traceId", traceId)

		m.poolConnection.Submit(func() {
			log.Message(
				traceId,
				"OUT",
				"GO-FIBER",
				"",
				"URL",
				ctx.OriginalURL(),
				"",
				"RESPONSE",
				string(ctx.Response().Body()))
		})
		return ctx.Next()
	}
}
