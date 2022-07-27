package app_utils

import (
	"fmt"
	"os"
)

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-sftp",
	Short: "A brief description of your application",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
