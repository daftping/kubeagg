package cmd

import (
	"github.com/daftping/kubeagg/pkg/kubeagg"
	"github.com/spf13/cobra"
)

var getConfig kubeagg.Config

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Kubectl wrapper to run against multiple contexts",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO Check for null
		getConfig.ObjectType = args[0]
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

	getCmd.PersistentFlags().StringVarP(
		&getConfig.Output,
		"output", "o",
		"table",
		"Output format. Supported values: table, wide, json.",
	)
	getCmd.PersistentFlags().StringVarP(
		&getConfig.Namespace,
		"namespace", "n",
		"default",
		"Namespace to operate in.",
	)
	getCmd.PersistentFlags().StringSliceVarP(
		&getConfig.Contexts,
		"contexts", "c",
		[]string{},
		"Send request to provided contexts. Has precedence over --context-pattern."+
			"(default: '', --context-pattern is used)",
	)
	getCmd.PersistentFlags().StringVarP(
		&getConfig.ContextPattern,
		"context-pattern", "p",
		".*",
		"Send request to contexts matched provided regexp. Ignored if --contexts is provided.",
	)
}
