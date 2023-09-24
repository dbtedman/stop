package main

import (
	"github.com/spf13/cobra"
)

func RunRoot(errorCh *chan error) {
	rootCommand := RootCommand(errorCh)

	err := rootCommand.Execute()

	if err != nil {
		*errorCh <- err
	} else {
		*errorCh <- nil
	}
}

func RootCommand(errorCh *chan error) *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "conveyance",
		Short: "Provide security by proxying requests to legacy applications.",
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()

			if err != nil {
				*errorCh <- err
			} else {
				*errorCh <- nil
			}
		},
	}

	var proxyServer Proxy = &ServerProxy{}

	rootCommand.AddCommand(ServeCommand(proxyServer, errorCh))
	rootCommand.AddCommand(VersionCommand(errorCh))

	return rootCommand
}
