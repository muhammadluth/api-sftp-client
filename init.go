package main

import (
	"fiber-demo-sftp/app/middleware"
	"fiber-demo-sftp/app/router"
	"fiber-demo-sftp/config"
	"fiber-demo-sftp/handler/mapper"
	"fiber-demo-sftp/handler/sftp_job"
	"fiber-demo-sftp/handler/usecase"

	"github.com/muhammadluth/log"
)

func RunningApplication() {
	properties := config.LoadConfig()
	log.SetupLogging(properties.LogPath)
	timeout := config.ParseTimeDuration(properties.Timeout)

	// SERVICE
	iFiberDemoSftpMapper := mapper.NewFiberDemoSftpMapper()

	iSftpInitiator := sftp_job.NewSftpInitiator(properties)
	iSftpRetrieveFile := sftp_job.NewSftpRetrieveFile(iSftpInitiator)
	iSftpSendFile := sftp_job.NewSftpSendFile(iSftpInitiator)
	iSftpDeleteFile := sftp_job.NewSftpDeleteFile(iSftpInitiator)

	iValidationUsecase := usecase.NewValidationUsecase()
	iRetrieveFileUsecase := usecase.NewRetrieveFileUsecase(iFiberDemoSftpMapper, iSftpRetrieveFile,
		iValidationUsecase)
	iSendFileUsecase := usecase.NewSendFileUsecase(iFiberDemoSftpMapper, iSftpSendFile,
		iValidationUsecase)
	iDeleteFileUsecase := usecase.NewDeleteFileUsecase(iSftpDeleteFile, iValidationUsecase)

	// GATEWAY
	iMiddleWare := middleware.NewMiddleware(properties)
	iSetupRouter := router.NewSetupRouter(timeout, properties, iMiddleWare, iRetrieveFileUsecase,
		iSendFileUsecase, iDeleteFileUsecase)

	iSetupRouter.Router()
}
