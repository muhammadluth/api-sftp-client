package sftp_job

import (
	"api-sftp-client/handler"
	"errors"
	"os"

	"github.com/muhammadluth/log"
)

type SftpRetrieveFile struct {
	iSftpInitiator handler.ISftpInitiator
}

func NewSftpRetrieveFile(iSftpInitiator handler.ISftpInitiator) handler.ISftpRetrieveFile {
	return &SftpRetrieveFile{iSftpInitiator}
}

func (s *SftpRetrieveFile) GetFile(traceId, pathDirectory string) (files []os.FileInfo, err error) {
	sftpClient, connection, err := s.iSftpInitiator.InitClientSftp(traceId)
	log.Event(traceId, "START", "Retrieve File")
	if err != nil {
		log.Error(err, traceId)
		return nil, errors.New("Error Connection To SFTP Server")
	}
	defer connection.Close()
	defer sftpClient.Close()

	if pathDirectory == "" {
		files, err = sftpClient.ReadDir(".")
		if err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Error Retrieve File From SFTP")
		}
	} else if pathDirectory != "" {
		if _, err := sftpClient.Lstat(pathDirectory); err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Directory Does Not Exist")
		}

		files, err = sftpClient.ReadDir(pathDirectory)
		if err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Error Retrieve File From SFTP")
		}
	}

	log.Event(traceId, "DONE", "Retrieve File")
	return files, nil
}

func (s *SftpRetrieveFile) GetDirectory(traceId, pathDirectory string) (directory []os.FileInfo, err error) {
	sftpClient, connection, err := s.iSftpInitiator.InitClientSftp(traceId)
	log.Event(traceId, "START", "Retrieve Directory")
	if err != nil {
		log.Error(err, traceId)
		return nil, errors.New("Error Connection To SFTP Server")
	}
	defer connection.Close()
	defer sftpClient.Close()

	if pathDirectory == "" {
		directory, err = sftpClient.ReadDir(".")
		if err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Error Retrieve Directory From SFTP")
		}
	} else if pathDirectory != "" {
		if _, err := sftpClient.Lstat(pathDirectory); err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Directory Does Not Exist")
		}

		directory, err = sftpClient.ReadDir(pathDirectory)
		if err != nil {
			log.Error(err, traceId)
			return nil, errors.New("Error Retrieve Directory From SFTP")
		}
	}

	log.Event(traceId, "DONE", "Retrieve Directory")
	return directory, nil
}
