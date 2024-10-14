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
	Short: "Start the local server that proxies requests for frontend and backend",
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
}
