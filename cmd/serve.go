package cmd

import (
	"github.com/spf13/cobra"
	"goauth/internal/api"
)

var portFlag int = 9888
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "starts the server",
	Long:  `starts the server on port default port 9888`,
	Run: func(cmd *cobra.Command, args []string) {
		api.Serve(portFlag)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&portFlag, "port", "p", portFlag, "port to listen on")

}


