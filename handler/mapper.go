package handler

import (
	"fiber-demo-sftp/model"
	"os"
)

type IFiberDemoSftpMapper interface {
	ToResponseGetFileFromSftp(pathDirectory string, files []os.FileInfo) (response []model.ResponseGetFileSFTP)
	ToResponseGetDirectoryFromSftp(directory []os.FileInfo) (response []model.ResponseGetDirectorySFTP)
}
