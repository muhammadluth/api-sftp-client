package router

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"api-sftp-client/app/middleware"
	"api-sftp-client/handler"
	"api-sftp-client/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/muhammadluth/log"
)

type SetupRouter struct {
	timeout              time.Duration
	properties           model.Properties
	iMiddleWare          middleware.Middleware
	iRetrieveFileUsecase handler.IRetrieveFileUsecase
	iSendFileUsecase     handler.ISendFileUsecase
	iDeleteFileUsecase   handler.IDeleteFileUsecase
}

func NewSetupRouter(timeout time.Duration, properties model.Properties,
	iMiddleWare middleware.Middleware, iRetrieveFileUsecase handler.IRetrieveFileUsecase,
	iSendFileUsecase handler.ISendFileUsecase, iDeleteFileUsecase handler.IDeleteFileUsecase) SetupRouter {
	return SetupRouter{timeout, properties, iMiddleWare, iRetrieveFileUsecase,
		iSendFileUsecase, iDeleteFileUsecase}
}

func (r *SetupRouter) Router() {
	addr := flag.String("addr", ":"+r.properties.Port, "TCP address to listen to")
	app := fiber.New()
	app.Use(etag.New())
	app.Use(compress.New())
	app.Use(requestid.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		ExposeHeaders:    "*",
		AllowMethods:     "GET, POST, PUT, DELETE",
	}))

	api := app.Group("/api/v1")
	api.Use("/temp", filesystem.New(filesystem.Config{
		Root:   http.Dir("./temp"),
		Browse: true,
		MaxAge: 3600,
	}))

	// RESTFULL API
	api.Get("/retrieve-file", r.iMiddleWare.ServiceMiddleware(), r.iRetrieveFileUsecase.GetFileSFTP)
	api.Get("/retrieve-directory", r.iMiddleWare.ServiceMiddleware(), r.iRetrieveFileUsecase.GetDirectorySFTP)
	api.Post("/send-file", r.iMiddleWare.ServiceMiddleware(), r.iSendFileUsecase.SendFileSFTP)
	api.Delete("/remove-file/*", r.iMiddleWare.ServiceMiddleware(), r.iDeleteFileUsecase.DeleteFileSFTP)
	api.Delete("/remove-directory", r.iMiddleWare.ServiceMiddleware(), r.iDeleteFileUsecase.DeleteDirectorySFTP)

	// HEALTH CHECK
	app.Get("/", r.iMiddleWare.ServiceMiddleware(), func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"message": "Hello, Welcome to My Apps!"})
	})

	log.Event("Listening on port" + *addr)
	fmt.Println("Listening on port" + *addr)
	fmt.Println("Ready to serve")
	log.Fatal(app.Listen(*addr))
}
