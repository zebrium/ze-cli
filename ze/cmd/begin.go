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

// beginCmd represents the begin command
var beginCmd = &cobra.Command{
	Use:   "begin",
	Short: "initialize a batch",
	Long:  `Initialize a Batch Upload to Zebrium.  This is the first step in submitting a batched bundle.`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ValidateAuthToken(viper.GetString("auth"))
		common.ValidateZapiUrl(viper.GetString("url"))
		if batchId != "" {
			batch.ValidateId(batchId)
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
	beginCmd.Flags().StringVarP(&batchId, "batchId", "b", "", "Sets custom batchId.  If not set, a random id will be generated")
}
