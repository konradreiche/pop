package cmd

import (
	"github.com/gobuffalo/pop"
	"github.com/spf13/cobra"
)

var all bool

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drops databases for you",
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if all {
			for env, conn := range pop.Connections {
				err = pop.DropDB(conn)
				if err != nil {
					pop.Logger.WithField("environment", env).WithError(err).Error("Failed to drop database")
				}
			}
		} else {
			if err := pop.DropDB(getConn()); err != nil {
				pop.Logger.WithError(err).Error("Failed to drop database")
			}
		}
	},
}

func init() {
	dropCmd.Flags().BoolVarP(&all, "all", "a", false, "Drops all of the databases in the database.yml")
	RootCmd.AddCommand(dropCmd)
}
