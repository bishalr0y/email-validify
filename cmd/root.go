package cmd

import (
	"fmt"
	"os"

	"github.com/bishalr0y/email-validify/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "email-validify",
	Short: "A cli tool to validate an email address",
}

func Execute() {
	fmt.Println(utils.GetBanner())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
