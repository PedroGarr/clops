/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strings"

	"github.com/PedroGarr/clops/utils"
	"github.com/spf13/cobra"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		// Host data treatment
		if host == "" {
			cmd.Println("Please inform the host to be validated")
			return
		}

		ports, _ := cmd.Flags().GetString("p")
		// Ports data treatment
		if ports == "" {
			cmd.Println("Please inform the ports / list of ports to be validated")
			return
		}

		// Return a slice of strings separated by commas and pass as parameter to the actual function
		portList := strings.Split(ports, ",")
		utils.CheckPort(host, portList)
	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)

	checkportCmd.PersistentFlags().String("h", "", "Host to be validated")
	checkportCmd.PersistentFlags().String("p", "", "Port list divided by comma. Ex: 80,443,8080")
}
