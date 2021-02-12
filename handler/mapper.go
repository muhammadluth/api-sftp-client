package handler

import (
	"api-sftp-client/model"
	"os"
)

type IApiSftpClientMapper interface {
	ToResponseGetFileFromSftp(pathDirectory string, files []os.FileInfo) (response []model.ResponseGetFileSFTP)
	ToResponseGetDirectoryFromSftp(directory []os.FileInfo) (response []model.ResponseGetDirectorySFTP)
}
