package cmd

import (
	"os"
	"path/filepath"

	"github.com/gobuffalo/pop"
	"github.com/markbates/going/defaults"
	"github.com/spf13/cobra"
)

var cfgFile string
var env string
var version bool

var RootCmd = &cobra.Command{
	Short: "A tasty treat for all your database needs",
	PersistentPreRun: func(c *cobra.Command, args []string) {
		pop.Logger.WithField("version", Version).Info("Starting")
		env = defaults.String(os.Getenv("GO_ENV"), env)
		setConfigLocation()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if !version {
			cmd.Help()
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		pop.Logger.WithError(err).Error("Failed")
	}
}

func init() {
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "Show version information")
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The configuration file you would like to use.")
	RootCmd.PersistentFlags().StringVarP(&env, "env", "e", "development", "The environment you want to run migrations against. Will use $GO_ENV if set.")
	RootCmd.PersistentFlags().BoolVarP(&pop.Debug, "debug", "d", false, "Use debug/verbose mode")
}

func setConfigLocation() {
	if cfgFile != "" {
		abs, err := filepath.Abs(cfgFile)
		if err != nil {
			return
		}
		dir, file := filepath.Split(abs)
		pop.AddLookupPaths(dir)
		pop.ConfigName = file
	}
	pop.LoadConfigFile()
}

func getConn() *pop.Connection {
	conn := pop.Connections[env]
	if conn == nil {
		pop.Logger.WithField("environment", env).Fatal("There is no connection for this environment")
	}
	return conn
}
