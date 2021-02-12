package sftp_job

import (
	"api-sftp-client/handler"
	"errors"

	"github.com/muhammadluth/log"
)

type SftpDeleteFile struct {
	iSftpInitiator handler.ISftpInitiator
}

func NewSftpDeleteFile(iSftpInitiator handler.ISftpInitiator) handler.ISftpDeleteFile {
	return &SftpDeleteFile{iSftpInitiator}
}

func (s *SftpDeleteFile) DeleteDirectory(traceId, pathDirectory string) (err error) {
	sftpClient, connection, err := s.iSftpInitiator.InitClientSftp(traceId)
	log.Event(traceId, "START", "Delete Directory")
	if err != nil {
		log.Error(err, traceId)
		return errors.New("Error Connection To SFTP Server")
	}
	defer connection.Close()
	defer sftpClient.Close()

	if _, err := sftpClient.Lstat(pathDirectory); err != nil {
		log.Error(err, traceId)
		return errors.New("Directory Does Not Exist")
	}

	err = sftpClient.RemoveDirectory(pathDirectory)
	if err != nil {
		log.Error(err, traceId)
		return errors.New("Error Delete Directory From SFTP")
	}

	log.Event(traceId, "DONE", "Delete Directory")
	return nil
}

func (s *SftpDeleteFile) DeleteFile(traceId, pathDirectory string) (err error) {
	sftpClient, connection, err := s.iSftpInitiator.InitClientSftp(traceId)
	log.Event(traceId, "START", "Delete File")
	if err != nil {
		log.Error(err, traceId)
		return errors.New("Error Connection To SFTP Server")
	}
	defer connection.Close()
	defer sftpClient.Close()

	if _, err := sftpClient.Lstat(pathDirectory); err != nil {
		log.Error(err, traceId)
		return errors.New("File or Directory Does Not Exist")
	}

	err = sftpClient.Remove(pathDirectory)
	if err != nil {
		log.Error(err, traceId)
		return errors.New("Error Delete File From SFTP")
	}

	log.Event(traceId, "DONE", "Delete File")
	return nil
}
