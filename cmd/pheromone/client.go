package pheromone

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var ClientCmd = &cobra.Command{
	Use:     "client [string to client]",
	Aliases: []string{"cli"},
	Short:   "client cmd",
	Long: `client.
this is client.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("client args: " + strings.Join(args, " "))
	},
}
