// Package cmd Copyright © 2023 ScienceLogic Inc/*
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// incidentsCmd represents the incidents command
var batchIncidentsCmd = &cobra.Command{
	Use:   "incidents",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("incidents called")
	},
}

func init() {
	batchCmd.AddCommand(batchIncidentsCmd)
	batchIncidentsCmd.Flags().String("timeFrom", "1", "Include Incidents created after this epoch time (use 1 as beginning of time)")
	batchIncidentsCmd.Flags().String("timeTo", "999999999999", "Include Incidents created before this epoch time (use 999999999999 as all time)")
	batchIncidentsCmd.Flags().String("timezone", "UTC", "Time zone name for time_from - time_to specification.")
	batchIncidentsCmd.Flags().String("repeatingIncidents", "first", "Include 'first' or 'all' occurrence(s) of an Incident Type")
	batchIncidentsCmd.Flags().StringP("batchId", "b", "", "Batch ID (required)")
	err := batchIncidentsCmd.MarkFlagRequired("batchId")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	batchIncidentsCmd.Hidden = true
}
