package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "store",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.WithError(err).Error("command failed")
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	_ = rootCmd.MarkPersistentFlagRequired("config")
}
