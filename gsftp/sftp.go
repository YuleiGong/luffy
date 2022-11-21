package gsftp

import "golang.org/x/crypto/ssh"

type Server struct {
	user           string
	password       string
	privateKeyPath string
	sshConfig      *ssh.ServerConfig
}

type ServerOpt func(*Server)

func WithUser(user, password string) ServerOpt {
	return func(s *Server) {
		s.user = user
		s.password = password
	}()
}

func NewServer(privateKey string, opts ...ServerOpt) *Server {
	s := &Server{
		privateKeyPath: privateKey,
	}
	for _, opt := range opts {
		opt(s)
	}
	s.sshConfig = newSSHConfig(s)

	return s
}

func (s *Server) Start() (err error) {

}
