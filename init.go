package main

import (
	"api-sftp-client/app/middleware"
	"api-sftp-client/app/router"
	"api-sftp-client/config"
	"api-sftp-client/handler/mapper"
	"api-sftp-client/handler/sftp_job"
	"api-sftp-client/handler/usecase"

	"github.com/muhammadluth/log"
)

func RunningApplication() {
	properties := config.LoadConfig()
	log.SetupLogging(properties.LogPath)
	timeout := config.ParseTimeDuration(properties.Timeout)

	// SERVICE
	iApiSftpClientMapper := mapper.NewApiSftpClientMapper()

	iSftpInitiator := sftp_job.NewSftpInitiator(properties)
	iSftpRetrieveFile := sftp_job.NewSftpRetrieveFile(iSftpInitiator)
	iSftpSendFile := sftp_job.NewSftpSendFile(iSftpInitiator)
	iSftpDeleteFile := sftp_job.NewSftpDeleteFile(iSftpInitiator)

	iValidationUsecase := usecase.NewValidationUsecase()
	iRetrieveFileUsecase := usecase.NewRetrieveFileUsecase(iApiSftpClientMapper, iSftpRetrieveFile,
		iValidationUsecase)
	iSendFileUsecase := usecase.NewSendFileUsecase(iApiSftpClientMapper, iSftpSendFile,
		iValidationUsecase)
	iDeleteFileUsecase := usecase.NewDeleteFileUsecase(iSftpDeleteFile, iValidationUsecase)

	// GATEWAY
	iMiddleWare := middleware.NewMiddleware(properties)
	iSetupRouter := router.NewSetupRouter(timeout, properties, iMiddleWare, iRetrieveFileUsecase,
		iSendFileUsecase, iDeleteFileUsecase)

	iSetupRouter.Router()
}
