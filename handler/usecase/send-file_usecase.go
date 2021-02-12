package usecase

import (
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model"
	"fiber-demo-sftp/model/constant"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type SendFileUsecase struct {
	iFiberDemoSftpMapper handler.IFiberDemoSftpMapper
	iSftpSendFile        handler.ISftpSendFile
	iValidationUsecase   handler.IValidationUsecase
}

func NewSendFileUsecase(iFiberDemoSftpMapper handler.IFiberDemoSftpMapper,
	iSftpSendFile handler.ISftpSendFile,
	iValidationUsecase handler.IValidationUsecase) handler.ISendFileUsecase {
	return &SendFileUsecase{iFiberDemoSftpMapper, iSftpSendFile, iValidationUsecase}
}

func (u *SendFileUsecase) SendFileSFTP(ctx *fiber.Ctx) error {
	var traceId = ctx.Locals("traceId").(string)

	pathDirectory := ctx.FormValue("path_directory")
	filename := ctx.FormValue("filename")
	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: "The File From The Request Encountered An Error",
		})
	}

	if isValid := u.iValidationUsecase.ValidatePathDirectory(traceId, pathDirectory); !isValid {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: "Invalid Directory SFTP",
		})
	}

	newFilename, err := u.iSftpSendFile.SendFile(traceId, pathDirectory, filename, file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	return ctx.JSON(model.ResponseHttp{
		Status:  constant.SUCCESS,
		Message: strings.Replace(constant.RESPONSE_SUCCESS_SEND, "{:file}", newFilename, 1),
	})
}
