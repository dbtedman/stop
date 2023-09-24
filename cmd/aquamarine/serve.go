package main

import (
	"github.com/spf13/cobra"
)

func ServeCommand(proxyServer Proxy, errorCh *chan error) *cobra.Command {
	var listenAddress string
	var proxyAddress string

	cmd := &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			theOptions, err := NewOptions(listenAddress, proxyAddress)

			if err != nil {
				*errorCh <- err
				return
			}

			ListenHTTPWithProxy(proxyServer, theOptions, errorCh)
		},
	}

	cmd.PersistentFlags().StringVar(&listenAddress, "from", ":3000", "")
	cmd.PersistentFlags().StringVar(&proxyAddress, "to", "", "")

	return cmd
}
