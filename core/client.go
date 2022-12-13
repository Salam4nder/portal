package core

import (
	"net"

	"github.com/Salam4nder/portal/config"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sshDialer interface {
	Dial(n, addr string) (net.Conn, error)
}

type sftpFileOpener interface {
	OpenFile(path string, f int) (*sftp.File, error)
}

type Client struct {
	config     *config.Config
	sshClient  sshDialer
	sftpClient sftpFileOpener
}

func NewClient(c *config.Config) *Client {
	return &Client{
		config:     c,
		sshClient:  &ssh.Client{},
		sftpClient: &sftp.Client{},
	}
}
