package main

import (
	"flag"
	"log"
	_ "net/http/pprof"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/unit"

	uart_server "github.com/taniho0707/HagoniwaMouse/server/uart"
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
				log.Printf("[UART] %v\n", string(buf))
				logCh <- string(buf)
			}
			// TODO: ReConnect
		}
	}()

	go func() {
		var w app.Window
		w.Option(app.Title("HagoniwaMouse"), app.Size(unit.Dp(1200), unit.Dp(800)))

		mainUI, err := mainApp.New(&w)
		if err != nil {
			log.Fatal(err)
		}
		mainUI.SetChannels(logCh)
		if err := mainUI.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()

	app.Main()
}
