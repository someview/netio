//go:build linux
// +build linux

package netio

import "syscall"

type EpollM struct {
	conn     map[int]*ServerConn
	socketFd int
	epollFd  int
}

type ServerConn struct{}

func (e *EpollM) Listen(ipAddr string, port int) error {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM)
	return nil
}
