package main

import (
	"github.com/dbtedman/stop/cmd/aquamarine/cmd"
	"github.com/dbtedman/stop/internal/signals"
)

func main() {
	signals.Bootstrap(run, performCleanup)
}

func run(errorCh *chan error) {
	cmd.RunRoot(errorCh)
}

func performCleanup(err error) {
	// Cleanup and resources used by this application on close.
}
