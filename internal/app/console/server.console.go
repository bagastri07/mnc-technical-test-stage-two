package console

import (
	"github.com/bagastri07/gin-boilerplate-service/internal/app/bootstrap"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command.
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a server process",
	Long: `The server command starts a server process that hosts and handles incoming requests from clients.
This command is used to launch a server application that listens for network connections and provides services or resources to clients.

Examples:

Start a web server to serve HTTP requests and deliver web pages or APIs.
Run a file server to share files over a local network or the internet.
`,
	Run: func(cmd *cobra.Command, args []string) {
		bootstrap.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
