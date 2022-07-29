package main

import (
	"github.com/bwoff11/frens/internal/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Debug("Starting frens")

	cli.Execute()
}
