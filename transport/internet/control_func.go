package internet

import "syscall"

type ControlFunc = func(network, address string, conn syscall.RawConn) error
