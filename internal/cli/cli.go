package cli

import (
	"fmt"
	"os"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/router"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frens",
	Short: "Frens server",
	Long:  `Frens is a federated social network server written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect()
		router.Execute()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
