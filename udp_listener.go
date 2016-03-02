package go_gps_tracker

import (
	"fmt"
	"net"
	//"log"
)

func Listener(port int, c chan Message, quit chan int) {
	ServerAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%d", port))
	error_check(err)

	ServerConn, err := net.ListenUDP("udp4", ServerAddr)
	error_check(err)
	defer ServerConn.Close()

	buf := make([]byte, 200)
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		error_check(err)
		m := Message{
			size:   n,
			msg:    string(buf[0:n]),
			source: addr.IP.String(),
		}
		//log.Printf("Received message [%s] from [%s]", m.msg, m.source)
		c <- m
	}
	close(c)
	quit <- 0
}
