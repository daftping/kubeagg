package cmd

import (
	"fmt"
	"os"

	"github.com/daftping/kubeagg/pkg/kubeagg"
	"github.com/spf13/cobra"
)

var GlobalConfig kubeagg.GlobalConfig

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubeagg",
	Short: "Kubectl wrapper to run against multiple contexts and namespaces",
	Long: `Kubectl wrapper can get any objects in any cluster in any namespace.
	
	You can provide list of contexts or context-pattern (regexp)
	If object you are trying to get namespaced you can provide list of 
	namespaces or namespace-pattern (regexp)
`,
	Example: `
	// Get namespaces from all contexts in (kubectl config view)
	kubeagg get ns

	// Get pods from contexts matched 'dev$|test$' regexp and in 
	// namespaces matched 'default|test|dev' regexp
	kubeagg \
		--context-pattern='dev$|test$' \
		--namespace-pattern='default|test|dev' \
		get pod
	
	// Get all nodes in "docker-desktop" and "test" contexts
	kubeagg --contexts=docker-desktop,test get no
	`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// 	kubeagg.SetGlobalConfig(GlobalConfig)
	// 		fmt.Println(GlobalConfig.LogLevel)
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")
	rootCmd.PersistentFlags().StringVarP(
		&GlobalConfig.LogLevel,
		"loglevel", "l",
		"Error",
		"Debug, Info, Warn, Error, Fatal",
	)

	rootCmd.PersistentFlags().StringVarP(
		&GlobalConfig.Output,
		"output", "o",
		"table",
		"Output format. Supported values: table, wide, json.",
	)

	rootCmd.PersistentFlags().StringSliceVar(
		&GlobalConfig.Contexts,
		"contexts",
		[]string{},
		"Send request to provided contexts. Has precedence over --context-pattern."+
			"(default: '', --context-pattern is used)",
	)
	rootCmd.PersistentFlags().StringVar(
		&GlobalConfig.ContextPattern,
		"context-pattern",
		".*",
		"Send request to contexts matched provided regexp. Ignored if --contexts is provided.",
	)

	rootCmd.PersistentFlags().BoolVar(
		&GlobalConfig.NoHeaders,
		"no-headers",
		false,
		"Skip headers in output",
	)

	rootCmd.PersistentFlags().StringSliceVarP(
		&GlobalConfig.Namespaces,
		"namespaces", "n",
		[]string{},
		"List namespaces to get objects from."+
			"(default: '', --namespace-pattern is used)",
	)

	rootCmd.PersistentFlags().StringVar(
		&GlobalConfig.NamespacePattern,
		"namespace-pattern",
		".*",
		"Get objects from namespaces matched provided regexp. Ignored if --namespaces is provided.",
	)

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}

	// 	// Search config in home directory with name ".newApp" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".newApp")
	// }

	// viper.AutomaticEnv() // read in environment variables that match

	// // If a config file is found, read it in.
	// if err := viper.ReadInConfig(); err == nil {
	// 	fmt.Println("Using config file:", viper.ConfigFileUsed())
	// }
}
