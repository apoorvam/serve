// A simple web server
package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"time"
)

var silent = flag.Bool("s", false, "Use this if you do not want to open browser.")
var address = flag.String("a", "localhost:4000", "Specifies the address where you want the file server to be served.")

func main() {
	flag.Parse()
	var addr string
	addr = *address
	if !strings.Contains(addr, ":") {
		addr += ":4000"
	}
	if strings.HasPrefix(addr, ":") {
		addr = "localhost" + addr
	}
	target := "./"
	if len(flag.Args()) > 0 {
		target = flag.Arg(0)
	}
	serve(addr, target)
}

func serve(addr, target string) {
	fs := http.FileServer(http.Dir(target))
	s := &http.Server{Addr: addr, Handler: fs, ReadTimeout: 10 * time.Second, WriteTimeout: 10 * time.Second}

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("Failed to bind to given address on port. %s : %s\n", addr, err.Error())
		os.Exit(1)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt)
	go func() {
		for _ = range stopChan {
			fmt.Println("Stopping the server...")
			listener.Close()
			fmt.Println("Closed.")
		}
	}()

	fmt.Printf("Serving files via %s.\nUse Ctrl+c to stop.\n", addr)

	if !*silent {
		command := exec.Command("open", fmt.Sprintf("http://%s", addr))
		command.Run()
	}

	s.Serve(listener)
}
