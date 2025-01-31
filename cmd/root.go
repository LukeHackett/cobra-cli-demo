package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/LukeHackett/cobra-cli-demo/internal/logging"
	"github.com/LukeHackett/cobra-cli-demo/internal/model"
	"github.com/LukeHackett/cobra-cli-demo/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "cobra-cli-demo",
	Version: "1.0.0",
	Short:   "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Set the application context
		config := model.NewCliConfig(*cmd.Flags(), *viper.GetViper())
		ctx := utils.SetConfig(cmd.Context(), config)
		ctx = utils.SetProfile(ctx, config.Profile)
		cmd.SetContext(ctx)

		// Configure the logging setup
		logging.ConfigureLogging(config.Debug)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.ExecuteContext(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Define the global commands
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra-cli-demo.yaml)")
	rootCmd.PersistentFlags().String("profile", "default", "Use a specific profile from your xerxes config file")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Help message for toggle")
	rootCmd.PersistentFlags().MarkHidden("debug")

	// Add the sub commands
	rootCmd.AddCommand(listUsersCmd)
	rootCmd.AddCommand(configSetCmd)
	rootCmd.AddCommand(configSetupCmd)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	logging.ConfigureLogging(debug)

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		slog.Info("Discovering the config file")
		// Find home directory.
		// home, err := os.UserHomeDir()
		// cobra.CheckErr(err)

		// Search config in home directory with name ".cobra-cli-demo" (without extension).
		viper.AddConfigPath("/Users/lha13/code/sky/spikes/cobra-cli-demo")
		viper.SetConfigType("toml")
		viper.SetConfigName("example-config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		slog.Info("Using config file: ", viper.ConfigFileUsed())
	} else {
		slog.Error("Unable to load config file: ", err)
		fmt.Fprintln(os.Stderr, "ERROR Unable to load config file:", err)
		os.Exit(1)
	}
}
