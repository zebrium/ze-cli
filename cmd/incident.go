// Package cmd Copyright © 2023 ScienceLogic Inc /*
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// incidentCmd represents the incident command
var incidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "A list of commands for interacting with incidents",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incident called")
	},
}

func init() {
	rootCmd.AddCommand(incidentCmd)
	incidentCmd.Hidden = true
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// incidentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// incidentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
