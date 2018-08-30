package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version string
	build   string
)

var rootCmd = &cobra.Command{
	Use: "hussar",
}

// Setup populates the version and build fields
func Setup(versionStr string, buildStr string) {
	version = versionStr
	build = buildStr
	rootCmd.Short = fmt.Sprintf("The Hussar programming language - %s (build %s)", version, build)
	rootCmd.SetVersionTemplate(fmt.Sprintf("%s (build %s)", version, build))
}

// Execute executes the commands
func Execute() {
	rootCmd.AddCommand(run, interactive, env)
	if err := rootCmd.Execute(); err != nil {
		log.WithError(err).Fatal()
	}
}

func init() {
	cobra.OnInitialize(initialize)

	// Global flags
	rootCmd.PersistentFlags().String("log.level", "warn", "one of debug, info, warn, error or fatal")
	rootCmd.PersistentFlags().String("log.format", "text", "one of text or json")

	// Flag binding
	viper.BindPFlags(rootCmd.PersistentFlags())
}

func initialize() {
	// Environment variables
	viper.SetEnvPrefix("hussar")
	viper.AutomaticEnv()

	// Configuration file
	viper.SetConfigName("hs-config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.hussar/")
	if err := viper.ReadInConfig(); err != nil {
		log.Info("No valid configuration file found")
	}
	lvl := viper.GetString("log.level")
	l, err := log.ParseLevel(lvl)
	if err != nil {
		log.WithField("level", lvl).Warn("Invalid log level, fallback to 'warn'")
	} else {
		log.SetLevel(l)
	}
	switch viper.GetString("log.format") {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
	case "text":
		log.SetFormatter(&log.TextFormatter{})
	}
}
