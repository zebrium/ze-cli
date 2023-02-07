// Package cmd Copyright © 2023 ScienceLogic Inc
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
		resp, err := batch.Show(viper.GetString("url"), viper.GetString("auth"), batchId)
		if err != nil {
			return err
		}
		if resp.Data == nil {
			return err
		} else {
			fmt.Fprintf(cmd.OutOrStdout(),"State for batch upload %s is %s\n", resp.Data[0].BatchID, resp.Data[0].State)
		}
		return nil
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
