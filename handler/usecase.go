package handler

import "github.com/gofiber/fiber/v2"

type IRetrieveFileUsecase interface {
	GetFileSFTP(ctx *fiber.Ctx) error
	GetDirectorySFTP(ctx *fiber.Ctx) error
}

type ISendFileUsecase interface {
	SendFileSFTP(ctx *fiber.Ctx) error
}

type IDeleteFileUsecase interface {
	DeleteDirectorySFTP(ctx *fiber.Ctx) error
	DeleteFileSFTP(ctx *fiber.Ctx) error
}

type IValidationUsecase interface {
	ValidatePathDirectory(traceId, workDirectory string) bool
}
