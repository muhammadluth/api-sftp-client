package usecase

import (
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model"
	"fiber-demo-sftp/model/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/muhammadluth/log"
)

type RetrieveFileUsecase struct {
	iFiberDemoSftpMapper handler.IFiberDemoSftpMapper
	iSftpRetrieveFile    handler.ISftpRetrieveFile
	iValidationUsecase   handler.IValidationUsecase
}

func NewRetrieveFileUsecase(iFiberDemoSftpMapper handler.IFiberDemoSftpMapper,
	iSftpRetrieveFile handler.ISftpRetrieveFile,
	iValidationUsecase handler.IValidationUsecase) handler.IRetrieveFileUsecase {
	return &RetrieveFileUsecase{iFiberDemoSftpMapper, iSftpRetrieveFile, iValidationUsecase}
}

func (u *RetrieveFileUsecase) GetFileSFTP(ctx *fiber.Ctx) error {
	var (
		traceId       = ctx.Locals("traceId").(string)
		pathDirectory = new(model.QueryParams)
	)
	if err := ctx.QueryParser(pathDirectory); err != nil {
		log.Error(err, traceId)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: "Invalid Path Directory File SFTP",
		})
	}

	if pathDirectory.PathDirectory != "" {
		if isValid := u.iValidationUsecase.ValidatePathDirectory(traceId,
			pathDirectory.PathDirectory); !isValid {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
				Status:  constant.ERROR,
				Message: "Invalid Directory SFTP",
			})
		}
	}

	dataFile, err := u.iSftpRetrieveFile.GetFile(traceId, pathDirectory.PathDirectory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	response := u.iFiberDemoSftpMapper.ToResponseGetFileFromSftp(pathDirectory.PathDirectory,
		dataFile)
	return ctx.JSON(model.ResponseSuccessWithData{
		TotalData: len(response),
		Data:      response,
	})
}

func (u *RetrieveFileUsecase) GetDirectorySFTP(ctx *fiber.Ctx) error {
	var (
		traceId       = ctx.Locals("traceId").(string)
		pathDirectory = new(model.QueryParams)
	)

	if err := ctx.QueryParser(pathDirectory); err != nil {
		log.Error(err, traceId)
		return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: "Invalid Path Directory SFTP",
		})
	}

	if pathDirectory.PathDirectory != "" {
		if isValid := u.iValidationUsecase.ValidatePathDirectory(traceId,
			pathDirectory.PathDirectory); !isValid {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.ResponseHttp{
				Status:  constant.ERROR,
				Message: "Invalid Directory SFTP",
			})
		}
	}

	dataDirectory, err := u.iSftpRetrieveFile.GetDirectory(traceId, pathDirectory.PathDirectory)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.ResponseHttp{
			Status:  constant.ERROR,
			Message: err.Error(),
		})
	}

	response := u.iFiberDemoSftpMapper.ToResponseGetDirectoryFromSftp(dataDirectory)
	return ctx.JSON(model.ResponseSuccessWithData{
		TotalData: len(response),
		Data:      response,
	})
}
