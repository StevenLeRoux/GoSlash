package main

import "github.com/spf13/cobra"

var (
	app = &cobra.Command{
		Use: "goslash",
		Run: func(cmd *cobra.Command, args []string) {
			runService(cmd)
		},
	}
)

func init() {
	app.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $PWD/config.yaml|toml|json)")
	app.PersistentFlags().StringVarP(&cfgPath, "configPath", "p", "", "config file path (default is $PWD)")
	app.PersistentFlags().IntVarP(&logLevel, "logLevel", "l", 0, "LogLevel")

}
