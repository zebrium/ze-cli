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

// endCmd represents the end command
var endCmd = &cobra.Command{
	Use:   "end",
	Short: "end batch and begin processing",
	Long:  `Signals that a batch is done uploading and that it can now be processed`,
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
		resp, err := batch.End(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			return err
		}
		if resp.Data == nil {
			return err
		} else {
			_, err := fmt.Fprintf(cmd.OutOrStdout(),"State for batch upload %s is now %s", resp.Data.BatchID, resp.Data.State)
			if err != nil {
				return err
			}
		}
		return nil
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
