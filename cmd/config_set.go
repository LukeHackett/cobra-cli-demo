package cmd

import (
	"errors"
	"fmt"
	"log/slog"

	"github.com/LukeHackett/cobra-cli-demo/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// listUsersCmd represents the listUsers command
var configSetCmd = &cobra.Command{
	Use:   "set-config",
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
	Run: func(cmd *cobra.Command, args []string) {
		profile := utils.GetProfile(cmd.Context())
		for k, v := 0, 1; k < len(args)/2; k, v = k+1, v+1 {
			key := profile + "." + args[k]
			value := args[v]
			viper.Set(key, value)
		}

		err := viper.WriteConfig()
		if err != nil {
			utils.Die(err, 1)
		} else {
			fmt.Println("Config file '" + viper.ConfigFileUsed() + "' was updated.")
			slog.Info("Config file '" + viper.ConfigFileUsed() + "' was updated.")
		}
	},
}

func init() {
	// Here you will define your flags and configuration settings.
}
