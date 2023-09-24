package signals

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

const ErrorResult = 1
const SuccessResult = 0

func Bootstrap(run func(errorCh *chan error), cleanup func(err error)) {
	signalsCh := make(chan os.Signal, 1)
	errorCh := make(chan error)
	var resultErr error

	signal.Notify(signalsCh, os.Interrupt, syscall.SIGTERM)

	defer func() {
		cleanup(resultErr)

		if resultErr != nil {
			fmt.Println(resultErr)
			os.Exit(ErrorResult)
		}

		os.Exit(SuccessResult)
	}()

	go func() {
		run(&errorCh)
	}()

	select {
	case <-signalsCh:
	case resultErr = <-errorCh:
	}
}
