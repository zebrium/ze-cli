// Package cmd Copyright © 2023 ScienceLogic Inc/*
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/batch"
	"github.com/zebrium/ze-cli/common"
	"os"
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "get current state",
	Long:  `Retrieves the current state of the batch.`,
	Run: func(cmd *cobra.Command, args []string) {
		batchId, err := cmd.Flags().GetString("batchId")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = common.ValidateAuthToken(viper.GetString("auth"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = common.ValidateZapiUrl(viper.GetString("url"))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = common.ValidateBatchId(batchId)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		resp, err := batch.Show(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if resp.Data == nil {
			fmt.Printf("Error: An error was returned by the server.\nError Code: %d\nError Message: %s\nRequest Status: %s\n", resp.Code, resp.Message, resp.Status)
			os.Exit(1)
		} else {
			fmt.Printf("State for batch upload %s is %s\n", resp.Data[0].BatchID, resp.Data[0].State)
		}
	},
}

func init() {
	batchCmd.AddCommand(stateCmd)
	stateCmd.Flags().StringP("batchId", "b", "", "Batch ID (required)")
	err := stateCmd.MarkFlagRequired("batchId")
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
