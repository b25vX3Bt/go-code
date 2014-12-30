package main

/*  base http://reiki4040.hatenablog.com/entry/2014/10/03/002239 */

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	pid := os.Getpid()
	fmt.Printf("pid=%d\n", pid)

	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	exit_chan := make(chan int)
	go func() {
		for {
			s := <-signal_chan
			switch s {
			// kill -SIGHUP XXXX
			case syscall.SIGHUP:
				fmt.Println("hungup")

				// kill -SIGINT XXXX or Ctrl+c
			case syscall.SIGINT:
				fmt.Println("Warikomi")

				// kill -SIGTERM XXXX
			case syscall.SIGTERM:
				fmt.Println("force stop")
				exit_chan <- 0

				// kill -SIGQUIT XXXX
			case syscall.SIGQUIT:
				fmt.Println("stop and core dump")
				exit_chan <- 0

			default:
				fmt.Println("Unknown signal.")
				exit_chan <- 1
			}
		}
	}()

	code := <-exit_chan
	os.Exit(code)
}
