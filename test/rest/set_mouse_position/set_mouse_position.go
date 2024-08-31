package main

import (
	"fmt"
	"math"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	xandy := float32(90 + 45)
	a := float32(270)
	xandyBits := math.Float32bits(xandy)
	aBits := math.Float32bits(a)

	data := []byte{0x00, 0x0D, 0x01}
	data = append(data, byte(xandyBits>>24), byte(xandyBits>>16), byte(xandyBits>>8), byte(xandyBits))
	data = append(data, byte(xandyBits>>24), byte(xandyBits>>16), byte(xandyBits>>8), byte(xandyBits))
	data = append(data, byte(aBits>>24), byte(aBits>>16), byte(aBits>>8), byte(aBits))

	_, err = conn.Write(data)
	if err != nil {
		panic(err)
	}

	fmt.Println("send complete -> (x, y, a) = (90 + 45, 90 + 45, 270)")
}
