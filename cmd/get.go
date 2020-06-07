package cmd

import (
	"github.com/daftping/kubeagg/pkg/kubeagg"
	"github.com/spf13/cobra"
)

var getConfig kubeagg.Config

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO Check for null
		getConfig.ObjectType = args[0]
		kubeagg.Run(getConfig)
	},
	// // TODO aliases
	// ValidArgs: []string{
	// 	"ns", "namespace",
	// 	"pod", "pods",
	// 	"deploy", "deployment",
	// },
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	getCmd.PersistentFlags().StringVarP(&getConfig.Output, "output", "o", "table", "Output format")
	getCmd.PersistentFlags().StringVarP(&getConfig.Namespace, "namespace", "n", "default", "Namespace to operate in")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
