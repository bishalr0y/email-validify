package cmd

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/bishalr0y/email-validify/utils"
	"github.com/spf13/cobra"
)

var (
	verifier = emailverifier.NewVerifier()
)

func validify(e string) {
	ret, err := verifier.Verify(e)
	if err != nil {
		fmt.Println("Verify email address failed\nError: ", err)
		return
	}
	if !ret.Syntax.Valid {
		fmt.Println("Error: email address syntax is invalid")
		return
	}

	fmt.Printf("Email: %v\n Username: %v\n Domain: %v\n Valid: %v\n Has MX Records: %v\n SMTP: %v\n Free: %v\n",
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
	Short: "pass the email(s) to validify",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No input email(s) provided, use -h for more details")
			return
		}
		fmt.Println(utils.GetBanner())
		emails := args
		for _, mail := range emails {
			validify(mail)
			fmt.Println("-----")
		}
	},
}

func init() {
	rootCmd.AddCommand(emailCmd)
}
