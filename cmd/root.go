/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"node-describe/constants"

	// _"node-describe/constants"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
func cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node-describe",
		Short: "node details in a user friendly manner",
		Long: `node details in kubectl, wont allow us to get details like, how much is left and all
	we will use this tool to get those infos`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			jokeTerm, _ := cmd.Flags().GetString("config")
			constants.SetCfgFile(jokeTerm)
			term := constants.GetCfgFile()
			fmt.Println("testing this root", term)
		},
	}
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cmd.PersistentFlags().String("config", "", "config file (default is $HOME/.node-describe.yaml)")
	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cm := cmd()
	cm.AddCommand(getCmd)

	err := cm.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

}

// html templating
