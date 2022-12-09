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

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a batch",
	Long:  `Cancel a batch job and ends all further processing`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ValidateAuthToken(viper.GetString("auth"))
		common.ValidateZapiUrl(viper.GetString("url"))
		batch.ValidateId(batchId)
		resp, err := batch.Cancel(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		if resp.Data == nil {
			fmt.Printf("Error: An error was returned by the server.\nError Code: %d\nError Message: %s\nRequest Status: %s\n", resp.Code, resp.Message, resp.Status)
			os.Exit(1)
		} else {
			fmt.Printf("State for batch upload %s is now %s\n", resp.Data.BatchID, resp.Data.State)
		}
	},
}

func init() {
	batchCmd.AddCommand(cancelCmd)
	cancelCmd.Flags().StringVarP(&batchId, "batchId", "b", "", "Batch ID (required)")
	cancelCmd.MarkFlagRequired("batchId")
}
