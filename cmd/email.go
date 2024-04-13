/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/spf13/cobra"
)

var (
	verifier = emailverifier.NewVerifier()
)

func validify(e string) {
	ret, err := verifier.Verify(e)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("email address syntax is invalid")
		return
	}

	fmt.Printf("email: %v, username: %v, domain: %v, valid: %v, has_mx_records: %v, smtp: %v, free: %v",
		ret.Email,
		ret.Syntax.Username,
		ret.Syntax.Domain,
		ret.Syntax.Valid,
		ret.HasMxRecords,
		ret.SMTP,
		ret.Free,
	)
}

// emailCmd represents the email command
var emailCmd = &cobra.Command{
	Use:   "email",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("email called")
		validify(args[0])
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// emailCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// emailCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
