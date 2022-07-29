package cli

import (
	"fmt"
	"os"

	"github.com/bwoff11/frens/internal/db"
	"github.com/bwoff11/frens/internal/router"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "frens",
	Short: "Frens server",
	Long:  `Frens is a federated social network server written in Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		db.Connect()
		router.Start()
	},
}

func Execute() {
	logrus.Debug("Starting CLI interpretation")
	if err := rootCmd.Execute(); err != nil {
		logrus.Debug("Error executing CLI root command")
		fmt.Println(err)
		os.Exit(1)
	}
}
