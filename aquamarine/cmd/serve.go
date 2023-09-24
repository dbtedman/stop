package cmd

import (
	"github.com/dbtedman/stop/aquamarine/internal/options"
	"github.com/dbtedman/stop/aquamarine/internal/proxy"
	"github.com/dbtedman/stop/aquamarine/web"
	"github.com/spf13/cobra"
)

func ServeCommand(proxyServer proxy.Proxy, errorCh *chan error) *cobra.Command {
	var listenAddress string
	var proxyAddress string

	cmd := &cobra.Command{
		Use: "serve",
		Run: func(cmd *cobra.Command, args []string) {
			theOptions, err := options.NewOptions(listenAddress, proxyAddress)

			if err != nil {
				*errorCh <- err
				return
			}

			web.ListenHTTPWithProxy(proxyServer, theOptions, errorCh)
		},
	}

	cmd.PersistentFlags().StringVar(&listenAddress, "from", ":3000", "")
	cmd.PersistentFlags().StringVar(&proxyAddress, "to", "", "")

	return cmd
}
