package uart_server

import (
	"io"

	"go.bug.st/serial"
)

const RWC_BUFFER_SIZE = 4096

type UartServer struct {
	mode *serial.Mode
	rwc  io.ReadWriteCloser
	port *serial.Port
	buf  []byte
}

func NewUartServer() *UartServer {
	mode := &serial.Mode{
		BaudRate: 921600,
		DataBits: 8,
		StopBits: serial.OneStopBit,
		Parity:   serial.NoParity,
	}
	return &UartServer{
		mode: mode,
		rwc:  nil,
		port: nil,
		buf:  make([]byte, RWC_BUFFER_SIZE),
	}
}

func (u *UartServer) Open(portName string) error {
	port, err := serial.Open(portName, u.mode)
	if err != nil {
		return err
	}
	u.port = &port
	return nil
}

func (u *UartServer) Close() error {
	if u.port != nil {
		err := (*u.port).Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UartServer) Read() ([]byte, error) {
	if u.port == nil {
		return nil, io.ErrClosedPipe
	}
	n, err := (*u.port).Read(u.buf)
	if err != nil {
		return nil, err
	}
	if n == 0 {
		u.Close()
		return nil, io.ErrClosedPipe
	}
	return u.buf[:n], nil
}

func (u *UartServer) Write(data []byte) error {
	if u.port == nil {
		return io.ErrClosedPipe
	}
	_, err := (*u.port).Write(data)
	if err != nil {
		return err
	}
	return nil
}
