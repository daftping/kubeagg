package cmd

import (
	"github.com/daftping/kubeagg/pkg/kubeagg"
	"github.com/spf13/cobra"
)

var getConfig kubeagg.Config

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Kubectl wrapper to run against multiple contexts and namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			getConfig.ObjectType = args[0]
		} else {
			getConfig.ObjectType = "all"
		}

		kubeagg.SetGlobalConfig(GlobalConfig)
		kubeagg.SetConfig(getConfig)
		kubeagg.Run()
	},
	// Looks like it doesn't make sense we can get all objects
	// ValidArgs: []string{
	// 	"ns", "namespace",
	// 	"pod", "pods",
	// 	"deploy", "deployment",
	// },
}

func init() {
	rootCmd.AddCommand(getCmd)
}
