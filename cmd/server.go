package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var serverCmd = &cobra.Command{
	Use:     "server [string to server]",
	Aliases: []string{"svr"},
	Short:   "server cmd",
	Long: `server.
this is server.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server args: " + strings.Join(args, " "))
	},
}
