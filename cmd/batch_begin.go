// Package cmd Copyright © 2023 ScienceLogic Inc
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/batch"
	"github.com/zebrium/ze-cli/common"
)

// beginCmd represents the begin batch command
var beginCmd = &cobra.Command{
	Use:   "begin",
	Short: "initialize a batch",
	Long:  `Initialize a Batch Upload to Zebrium.  This is the first step in submitting a batched bundle.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		batchId, err := cmd.Flags().GetString("batchId")
		if err != nil {
			return err
		}
		err = common.ValidateAuthToken(viper.GetString("auth"))
		if err != nil {
			return err
		}
		err = common.ValidateZapiUrl(viper.GetString("url"))
		if err != nil {
			return err
		}
		if len(batchId) != 0 {
			err = common.ValidateBatchId(batchId)
			if err != nil {
				return err
			}
		}
		resp, err := batch.Begin(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			return err
		}
		if resp.Data == nil || len(resp.Data.BatchId) == 0 {
			return err
		} else {
			fmt.Fprintf(cmd.OutOrStdout(),"Batch Id: %s Successfully created.\n", resp.Data.BatchId)
		}
		return nil
	},
}

func init() {
	batchCmd.AddCommand(beginCmd)
	beginCmd.Flags().StringP("batchId", "b", "", "Sets custom batchId.  If not set, a random id will be generated")

}
