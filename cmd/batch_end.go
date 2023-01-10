// Package cmd Copyright © 2023 ScienceLogic Inc/*
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/batch"
	"github.com/zebrium/ze-cli/common"
	"os"

	"github.com/spf13/cobra"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "end batch and begin processing",
	Long:  `Signals that a batch is done uploading and that it can now be processed`,
	Run: func(cmd *cobra.Command, args []string) {
		batchId, err := cmd.Flags().GetString("batchId")
		if err != nil {
			fmt.Println(err)
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
		resp, err := batch.End(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Printf("Client: Error occured generating request. Error: %s\n", err.Error())
			os.Exit(1)
		}
		if resp.Data == nil {
			fmt.Printf("Error: An error was returned by the server.\nError Code: %d \nError Message: %s\nRequest Status: %s\n", resp.Code, resp.Message, resp.Status)
			os.Exit(1)
		} else {
			fmt.Printf("State for batch upload %s is now %s", resp.Data.BatchID, resp.Data.State)
		}
	},
}

func init() {
	batchCmd.AddCommand(endCmd)
	endCmd.Flags().StringP("batchId", "b", "", "Batch ID (required)")
	err := endCmd.MarkFlagRequired("batchId")
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
