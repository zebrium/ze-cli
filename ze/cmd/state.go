/*
Copyright © 2022 ScienceLogic Inc
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/ze/batch"
	"github.com/zebrium/ze-cli/ze/common"
	"os"
)

// stateCmd represents the state command
var stateCmd = &cobra.Command{
	Use:   "state",
	Short: "get current state",
	Long:  `Retrieves the current state of the batch.`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ValidateAuthToken(viper.GetString("auth"))
		common.ValidateZapiUrl(viper.GetString("url"))
		batch.ValidateId(batchId)
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
	stateCmd.Flags().StringVarP(&batchId, "batchId", "b", "", "Batch ID (required)")
	stateCmd.MarkFlagRequired("batchId")
}
