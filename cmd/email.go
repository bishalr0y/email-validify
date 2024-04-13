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
		fmt.Println("Verify email address failed, error is: ", err)
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
	Short: "pass the email to validify it",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Error: No input email provided, use -h for more details")
			return
		}
		text := `
		┌────────────────────────────────────────────────────────────────────────────────────────────────────┐
		│ _______ _______ _______ _____             _    _ _______        _____ ______  _____ _______ __   __│
		│ |______ |  |  | |_____|   |   |            \  /  |_____| |        |   |     \   |   |______   \_/  │
		│ |______ |  |  | |     | __|__ |_____        \/   |     | |_____ __|__ |_____/ __|__ |          |   │
		│                                                                                                    │
		└────────────────────────────────────────────────────────────────────────────────────────────────────┘
		`

		fmt.Println(text)
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
