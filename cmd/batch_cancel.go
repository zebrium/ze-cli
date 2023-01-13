// Package cmd Copyright © 2023 ScienceLogic Inc
package cmd

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/batch"
	"github.com/zebrium/ze-cli/common"
	"os"

	"github.com/spf13/cobra"
)

// cancelCmd represents the cancel command
var cancelCmd = &cobra.Command{
	Use:   "cancel",
	Short: "cancel a batch",
	Long:  `Cancel a batch job and ends all further processing`,
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
		err = common.ValidateBatchId(batchId)
		if err != nil {
			return err
		}
		resp, err := batch.Cancel(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			return err
		}
		if resp.Data == nil {
			return err
		} else {
			fmt.Fprintf(cmd.OutOrStdout(),"State for batch upload %s is now %s\n", resp.Data.BatchID, resp.Data.State)
		}
		return nil
	},
}

func init() {
	batchCmd.AddCommand(cancelCmd)
	cancelCmd.Flags().StringP("batchId", "b", "", "Batch ID (required)")
	err := cancelCmd.MarkFlagRequired("batchId")
	if err != nil {
		println(err)
		os.Exit(1)
	}
	err = viper.BindPFlags(cancelCmd.Flags())
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
