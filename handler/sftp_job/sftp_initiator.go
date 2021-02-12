package sftp_job

import (
	"fiber-demo-sftp/handler"
	"fiber-demo-sftp/model"

	"github.com/muhammadluth/log"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type SftpInitiator struct {
	properties model.Properties
}

func NewSftpInitiator(properties model.Properties) handler.ISftpInitiator {
	return &SftpInitiator{properties}
}

func (s *SftpInitiator) InitClientSftp(traceId string) (sftpClient *sftp.Client, connection *ssh.Client,
	err error) {
	hostSftp := s.properties.SftpConfig.Host
	configSftp := ssh.ClientConfig{
		User:            s.properties.SftpConfig.User,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(s.properties.SftpConfig.Password),
		},
	}
	connection, err = ssh.Dial("tcp", hostSftp, &configSftp)
	if err != nil {
		log.Error(err, traceId)
		return sftpClient, connection, err
	}
	log.Event(traceId, "INFO", "Connected to Remote Host")
	sftpClient, err = sftp.NewClient(connection)
	if err != nil {
		log.Error(err, traceId)
		return sftpClient, connection, err
	}
	return sftpClient, connection, err
}
