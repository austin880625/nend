/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/austin880625/nend/cmd_args"
	"github.com/austin880625/nend/server"
	"github.com/spf13/cobra"
)

var cmdArgs = cmd_args.Args{}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Run(cmdArgs)
	},
}

func init() {
	serveCmd.Flags().IntVarP(&cmdArgs.Port, "port", "p", 8000, "listen port")
	serveCmd.Flags().IntVarP(&cmdArgs.BPort, "bport", "b", 8080, "backend port")
	serveCmd.Flags().IntVarP(&cmdArgs.FPort, "fport", "f", 3000, "frontend port")
	serveCmd.Flags().StringVarP(&cmdArgs.ApiPath, "path", "a", "/api/v1", "api path prefix")
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
