package simplenet

import (
	"os"
	"syscall"
)

type netSocket struct {
	fd int
}

func NewNetSocket(ip IP, port int) (*netSocket, error) {
	syscall.ForkLock.Lock()
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, os.NewSyscallError("socket", err)
	}
	syscall.ForkLock.Unlock()
	if err = syscall.SetsockoptInt(fd, syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
		syscall.Close(fd)
		return nil, os.NewSyscallError("setsockopt", err)
	}
	// Bind the socket to a port
	sa := &syscall.SockaddrInet4{Port: port}
	copy(sa.Addr[:], ip)
	if err = syscall.Bind(fd, sa); err != nil {
		return nil, os.NewSyscallError("bind", err)
	}
	// Listen for incoming connections
	if err = syscall.Listen(fd, syscall.SOMAXCONN); err != nil {
		return nil, os.NewSyscallError("listen", err)
	}
	return &netSocket{fd: fd}, nil
}

func (ns netSocket) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	n, err := syscall.Read(ns.fd, p)
	if err != nil {
		n = 0
	}
	return n, err
}
