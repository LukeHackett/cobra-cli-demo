package cmd

import (
	"fmt"
	"log/slog"

	"github.com/LukeHackett/cobra-cli-demo/internal/service"
	"github.com/spf13/cobra"
)

// listUsersCmd represents the listUsers command
var listUsersCmd = &cobra.Command{
	Use:   "list-users",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listUsers called")
		slog.Debug("A debug Log")
		slog.Info("A info Log")
		slog.Warn("A warn Log")
		slog.Error("An error Log")

		svc := &service.UserService{BaseUrl: "https://jsonplaceholder.typicode.com"}

		res := svc.FetchAll()

		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(listUsersCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listUsersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listUsersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
