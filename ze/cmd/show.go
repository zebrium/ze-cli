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

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "get batch metrics",
	Long:  `Gets the current status and metrics for a given batch`,
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
			fmt.Println(resp.Data[0].String())
		}
	},
}

func init() {
	batchCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&batchId, "batchId", "b", "", "Batch ID (required)")
	showCmd.MarkFlagRequired("batchId")
}
