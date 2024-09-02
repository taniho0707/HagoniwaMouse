package udp_server

import (
	"errors"
	"fmt"
	"log"
	"net"

	udp_domain "github.com/taniho0707/HagoniwaMouse/server/domain"
	"github.com/taniho0707/HagoniwaMouse/simulator"
)

const RWC_BUFFER_SIZE = 4096

type UdpServer struct {
	conn     *net.UDPConn
	lastaddr *net.UDPAddr
	buf      []byte
}

func NewUdpServer() *UdpServer {
	return &UdpServer{
		conn: nil,
		buf:  make([]byte, RWC_BUFFER_SIZE),
	}
}

func (u *UdpServer) Open(port string) error {
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		return err
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return err
	}

	u.conn = conn
	return nil
}

func (u *UdpServer) Close() error {
	if u.conn != nil {
		err := u.conn.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *UdpServer) Read() ([]byte, *net.UDPAddr, error) {
	if u.conn == nil {
		return nil, nil, net.ErrClosed
	}
	n, addr, err := u.conn.ReadFromUDP(u.buf)
	if err != nil {
		return nil, nil, err
	}
	u.lastaddr = addr
	return u.buf[:n], addr, nil
}

func (u *UdpServer) Write(data []byte) error {
	if u.conn == nil {
		return net.ErrClosed
	}
	_, err := u.conn.WriteToUDP(data, u.lastaddr)
	if err != nil {
		return err
	}
	return nil
}

func (u *UdpServer) Run(udpReceiveCh chan udp_domain.UdpCommand, udpResponseCh chan udp_domain.UdpCommand) error {
	sim := simulator.NewSimulator()

	for {
		buf, addr, err := u.Read()
		if err != nil {
			return err
		}

		cmd, err := ParseUdpCommand(buf)
		if err != nil {
			log.Printf("[UDP] %v", err)
			continue
		}

		fmt.Printf("[UDP] %v %v\n", addr, cmd)

		// Simulator に値を渡して UdpCommand を上書き
		err = sim.Next(&cmd)
		if err != nil {
			log.Printf("[UDP Sim] %v", err)
			continue
		}

		// UdpCommand.Code が GetMousePosition のとき、100 回に 1 回 History に追加

		udpReceiveCh <- cmd
		v, ok := <-udpResponseCh
		if !ok {
			return errors.New("udpResponseCh is closed")
		}

		if v.Code != udp_domain.CommandInternalNoResponse {
			res, err := BuildBytesFromUdpCommand(&v)
			if err != nil {
				log.Printf("[UDP] %v", err)
				continue
			}

			err = u.Write(res)
			if err != nil {
				log.Printf("[UDP] %v", err)
				return err
			}
		}
	}
}
