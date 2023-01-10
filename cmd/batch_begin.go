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

// beginCmd represents the begin command
var beginCmd = &cobra.Command{
	Use:   "begin",
	Short: "initialize a batch",
	Long:  `Initialize a Batch Upload to Zebrium.  This is the first step in submitting a batched bundle.`,
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
		if batchId != "" {
			err = common.ValidateBatchId(batchId)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		resp, err := batch.Begin(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Printf("Client: Error occured generating request. Error: %s\n", err.Error())
			os.Exit(1)
		}
		if resp.Data == nil || resp.Data.BatchId == "" {
			fmt.Printf("Error: no batch id was returned by the server.\nError Code: %d\nError Message: %s\nRequest Status: %s\n", resp.Code, resp.Message, resp.Status)
			os.Exit(1)
		} else {
			fmt.Printf("Batch Id: %s Successfully created.\n", resp.Data.BatchId)
		}
	},
}

func init() {
	batchCmd.AddCommand(beginCmd)
	beginCmd.Flags().StringP("batchId", "b", "", "Sets custom batchId.  If not set, a random id will be generated")

}
