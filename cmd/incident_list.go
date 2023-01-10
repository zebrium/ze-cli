// Package cmd Copyright © 2023 ScienceLogic Inc/*
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/common"
	"github.com/zebrium/ze-cli/incidents"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all incidents in a given time range",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := common.ValidateZapiUrl(viper.GetString("url"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = common.ValidateAPIToken(viper.GetString("api"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp, err := incidents.List(viper.GetString("url"), viper.GetString("api"), viper.GetInt("timeFrom"), viper.GetInt("timeTo"), viper.GetString("repeatingIncidents"), viper.GetString("timezone"), viper.GetString("batchId"))
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if resp.Data == nil {
			fmt.Printf("Error: An error was returned by the server.\nError Code: %d\nError Message: %s\nRequest Status: %s\n", resp.Error.Code, resp.Error.Message, resp.Error.Data)
			os.Exit(1)
		} else {
			fmt.Println(resp.Data)
		}
	},
}

func init() {
	incidentCmd.AddCommand(listCmd)

	listCmd.Flags().Int("timeFrom", 1, "Include Incidents created after this epoch time (use 1 as beginning of time)")
	listCmd.Flags().Int("timeTo", 999999999999, "Include Incidents created before this epoch time (use 999999999999 as all time)")
	listCmd.Flags().String("timezone", "UTC", "Time zone name for time_from - time_to specification.")
	listCmd.Flags().String("repeatingIncidents", "first", "Include 'first' or 'all' occurrence(s) of an Incident Type")
	err := viper.BindPFlags(listCmd.Flags())
	if err != nil {
		println(err)
		os.Exit(1)
	}
	listCmd.Hidden = true
}
