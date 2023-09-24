package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const ErrorResult = 1
const SuccessResult = 0

func main() {
	signalsCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	var resultErr error

	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		if resultErr != nil {
			fmt.Println(resultErr)
			os.Exit(ErrorResult)
		}

		fmt.Println("\nThat's all done.")
		os.Exit(SuccessResult)
	}()

	go func() {
		RunRoot(&errorCh)
	}()

	select {
	case <-signalsCh:
	case resultErr = <-errorCh:
	}
}
