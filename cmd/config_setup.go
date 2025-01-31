package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/LukeHackett/cobra-cli-demo/internal/utils"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// listUsersCmd represents the listUsers command
var configSetupCmd = &cobra.Command{
	Use:   "setup-config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args)%2 != 0 {
			return errors.New("an even number of arguments should be supplied in the format key1 value1 key2 value2")
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if !term.IsTerminal(int(syscall.Stdin)) {
			return errors.New("terminal is not interactive! consider using the configure set commands")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		profile := utils.GetProfile(cmd.Context())

		baseUrl := StringPrompt("Enter the base url: ")
		secret := PasswordPrompt("Enter the client id: ")

		fmt.Println(profile)
		fmt.Println(baseUrl)
		fmt.Println(secret)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}

// PasswordPrompt asks for a string value using the label.
// The entered value will not be displayed on the screen
// while typing.
func PasswordPrompt(label string) string {
	var s string
	for {
		fmt.Fprint(os.Stderr, label+" ")
		b, _ := term.ReadPassword(int(syscall.Stdin))
		s = string(b)
		if s != "" {
			break
		}
	}
	fmt.Println()
	return s
}
