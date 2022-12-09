/*
Copyright © 2022 ScienceLogic Inc
*/
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/ze/batch"
	"github.com/zebrium/ze-cli/ze/common"
	"os"

	"github.com/spf13/cobra"
)

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "end batch and begin processing",
	Long:  `Signals that a batch is done uploading and that it can now be processed`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ValidateAuthToken(viper.GetString("auth"))
		common.ValidateZapiUrl(viper.GetString("url"))
		batch.ValidateId(batchId)
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
	endCmd.Flags().StringVarP(&batchId, "batchId", "b", "", "Batch ID (required)")
	endCmd.MarkFlagRequired("batchId")
}
