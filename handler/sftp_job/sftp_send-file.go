package sftp_job

import (
	"errors"
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model/constant"
	"io"
	"mime/multipart"
	"strings"

	"github.com/muhammadluth/log"
)

type SftpSendFile struct {
	iSftpInitiator handler.ISftpInitiator
}

func NewSftpSendFile(iSftpInitiator handler.ISftpInitiator) handler.ISftpSendFile {
	return &SftpSendFile{iSftpInitiator}
}

func (s *SftpSendFile) SendFile(traceId, pathDirectory, filename string,
	file *multipart.FileHeader) (string, error) {
	sftpClient, connection, err := s.iSftpInitiator.InitClientSftp(traceId)
	log.Event(traceId, "START", "Send File")
	if err != nil {
		log.Error(err, traceId)
		return "", errors.New("Error Connection To SFTP Server")
	}
	defer connection.Close()
	defer sftpClient.Close()

	if _, err := sftpClient.Lstat(pathDirectory); err != nil {
		log.Error(err, traceId)
		sftpClient.MkdirAll(pathDirectory)
	}

	splitFilename := strings.Split(file.Filename, ".")
	extensionFile := splitFilename[len(splitFilename)-1]
	newFilename := strings.Replace(constant.FORMAT_FILENAME, "{:filename}", filename, 1)
	newFilename = strings.Replace(newFilename, "{:extension}", extensionFile, 1)
	fileWithPath := pathDirectory + "/" + newFilename

	createFileInSftp, err := sftpClient.Create(fileWithPath)
	if err != nil {
		log.Error(err, traceId)
		return "", errors.New("Error Create File In SFTP Server")
	}

	openSourceFile, err := file.Open()
	if err != nil {
		log.Error(err, traceId)
		return "", errors.New("The File From The Request Encountered An Error")
	}
	_, err = io.Copy(createFileInSftp, openSourceFile)
	if err != nil {
		log.Error(err, traceId)
		return "", errors.New("Error Send File To SFTP")
	}

	log.Event(traceId, "DONE", "Send File")
	return newFilename, nil
}
