package handler

import (
	"mime/multipart"
	"os"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type ISftpInitiator interface {
	InitClientSftp(traceId string) (sftpClient *sftp.Client, connection *ssh.Client, err error)
}

type ISftpRetrieveFile interface {
	GetFile(traceId, pathDirectory string) (files []os.FileInfo, err error)
	GetDirectory(traceId, pathDirectory string) (directory []os.FileInfo, err error)
}

type ISftpSendFile interface {
	SendFile(traceId, pathDirectory, filename string, file *multipart.FileHeader) (string, error)
}

type ISftpDeleteFile interface {
	DeleteDirectory(traceId, pathDirectory string) (err error)
	DeleteFile(traceId, pathDirectory string) (err error)
}
