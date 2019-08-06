package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan)

	go func() {
		for {
			s := <-signalChan
			switch s {
			case syscall.SIGHUP:
				fmt.Println("Detected SIGHUP")

			case syscall.SIGINT:
				fmt.Println("Detected SIGINT")

			case syscall.SIGTERM:
				fmt.Println("Detected SIGTERM")

			case syscall.SIGQUIT:
				fmt.Println("Detected SIGQUIT")

			default:
				fmt.Println("Unknown signal")
			}
		}
	}()

	argc := len(os.Args)

	var n int
	var err error
	if argc < 2 {
		n = 1
	} else {
		n, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	var exitCode int
	if argc < 3 {
		exitCode = 0
	} else {
		exitCode, err = strconv.Atoi(os.Args[2])
		if err != nil {
			panic(err)
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Running: %d\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Finished")
	os.Exit(exitCode)
}
