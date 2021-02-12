package usecase

import (
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model"
	"fiber-demo-sftp/model/constant"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type DeleteFileUsecase struct {
	iSftpDeleteFile    handler.ISftpDeleteFile
	iValidationUsecase handler.IValidationUsecase
}

func NewDeleteFileUsecase(iSftpDeleteFile handler.ISftpDeleteFile,
	iValidationUsecase handler.IValidationUsecase) handler.IDeleteFileUsecase {
	return &DeleteFileUsecase{iSftpDeleteFile, iValidationUsecase}
}

func (u *DeleteFileUsecase) DeleteFileSFTP(ctx *fiber.Ctx) error {
	var (
		traceId       = ctx.Locals("traceId").(string)
		pathDirectory = ctx.Params("*")
	)

	if pathDirectory != "" {
		if isValid := u.iValidationUsecase.ValidatePathDirectory(traceId, pathDirectory); !isValid {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
				Status:  constant.ERROR,
				Message: "Invalid Directory SFTP",
			})
		}
	}

	err := u.iSftpDeleteFile.DeleteFile(traceId, pathDirectory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	name := strings.Split(pathDirectory, "/")
	return ctx.JSON(model.ResponseHttp{
		Status:  constant.SUCCESS,
		Message: strings.Replace(constant.RESPONSE_SUCCESS_DELETE, "{:file}", name[len(name)-1], 1),
	})
}

func (u *DeleteFileUsecase) DeleteDirectorySFTP(ctx *fiber.Ctx) error {
	return ctx.JSON(model.ResponseHttp{
		Status:  constant.SUCCESS,
		Message: "Sorry, This Feature Is Currently Under Development!",
	})
}
