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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "get batch metrics",
	Long:  `Gets the current status and metrics for a given batch`,
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
			output, err := resp.Data[0].String()
			if err != nil {
				return err
			}
			_, err = fmt.Fprintln(cmd.OutOrStdout(), output)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	batchCmd.AddCommand(showCmd)
	showCmd.Flags().StringP("batchId", "b", "", "Batch ID (required)")
	err := showCmd.MarkFlagRequired("batchId")
	if err != nil {
		println(err)
		os.Exit(1)
	}
}
