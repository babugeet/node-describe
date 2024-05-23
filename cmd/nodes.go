/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"node-describe/constants"
	"node-describe/internal/nodes"

	"github.com/spf13/cobra"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// if jokeTerm == "" {
		// 	fmt.Println("term is empty ")
		// } else {
		// 	fmt.Println("input recieved is ", jokeTerm)
		// }
		// fmt.Println(rootCmd.Flags().GetString("config"))
		jokeTerm, _ := cmd.Flags().GetString("config")
		constants.SetCfgFile(jokeTerm)
		// _ := constants.GetCfgFile()
		// fmt.Println("testing this nodes ", term)

		// fmt.Println("nodes called")
		// Create a new tab writer

		nodes.DescribeNode()
		// for _, j := range nodes.GetNodes() {
		// 	fmt.Println(j.Name)
		// }

	},
}

func init() {
	// rootCmd.AddCommand(nodesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
