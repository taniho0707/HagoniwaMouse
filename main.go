package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/unit"

	uart_server "github.com/taniho0707/HagoniwaMouse/server/uart"
	udp_server "github.com/taniho0707/HagoniwaMouse/server/udp"
	mainApp "github.com/taniho0707/HagoniwaMouse/ui/app"
)

// var (
// 	enablePprof = flag.Bool("pprof", false, "enable pprof")
// )

func main() {
	flag.Parse()

	// if *enablePprof {
	// 	go func() {
	// 		log.Println(http.ListenAndServe("localhost:6060", nil))
	// 	}()
	// }

	logCh := make(chan string)
	uartServer := uart_server.NewUartServer()
	go func() {
		var bufLine []byte
		for {
			if err := uartServer.Open("/dev/ttyACM0"); err != nil {
				time.Sleep(10 * time.Second)
				continue
			}
			for {
				buf, err := uartServer.Read()
				if err != nil {
					panic(err)
				}

				// 改行があればそこまででログ出力
				foundReturn := false
				for i, b := range buf {
					if b == '\n' {
						foundReturn = true
						bufLine = append(bufLine, buf[:i]...)
						fmt.Printf("[UART] %v", string(bufLine))
						logCh <- string(bufLine)
						if i+1 == len(buf) {
							bufLine = []byte{}
						} else {
							bufLine = buf[i+1:]
						}
						break
					}
				}
				// 改行がなければバッファに追加のみ
				if !foundReturn {
					bufLine = append(bufLine, buf...)
				}
			}
			// TODO: ReConnect
		}
	}()

	udpReceiveCh := make(chan udp_server.UdpCommand)
	udpResponseCh := make(chan udp_server.UdpCommand)
	udpServer := udp_server.NewUdpServer()
	go func() {
		if err := udpServer.Open(":3000"); err != nil {
			panic(err)
		}
		udpServer.Run(udpReceiveCh, udpResponseCh)
	}()

	go func() {
		var w app.Window
		w.Option(app.Title("HagoniwaMouse"), app.Size(unit.Dp(1200), unit.Dp(800)))

		mainUI, err := mainApp.New(&w)
		if err != nil {
			log.Fatal(err)
		}
		mainUI.SetChannels(logCh, udpReceiveCh, udpResponseCh)
		if err := mainUI.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
