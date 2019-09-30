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
				os.Exit(1)

			case syscall.SIGINT:
				fmt.Println("Detected SIGINT")
				os.Exit(1)

			case syscall.SIGTERM:
				fmt.Println("Detected SIGTERM")
				os.Exit(1)

			case syscall.SIGQUIT:
				fmt.Println("Detected SIGQUIT")
				os.Exit(1)

			default:
				fmt.Println("Unknown signal")
			}
		}
	}()

	n, err := strconv.Atoi(os.Getenv("TIMES"))
	if err != nil {
		panic(err)
	}
	exitStatus, exitErr := strconv.Atoi(os.Getenv("EXIT_STATUS"))
	if exitErr != nil {
		panic(err)
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("Running: %d\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("Finished")
	os.Exit(exitStatus)
}
