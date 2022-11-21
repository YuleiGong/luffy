package gsftp

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func newSSHConfig(s *Server) *ssh.ServerConfig {
	return &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			if c.User() == s.user && string(pass) == s.password {
				return nil, nil
			}
			return nil, fmt.Errorf("password rejected for %q", c.User())
		},
	}
}
