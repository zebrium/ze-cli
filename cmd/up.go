// Package cmd Copyright © 2023 ScienceLogic Inc
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/common"
	"github.com/zebrium/ze-cli/up"
	"os"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Upload a file to Zebrium",
	Long:  `Uploads a file to Zebrium for analysis`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := common.ValidateAuthToken(viper.GetString("auth"))
		if err != nil {
			return err
		}
		err = common.ValidateZapiUrl(viper.GetString("url"))
		if err != nil {
			return err
		}
		err = common.ValidateUpMetadata(viper.GetString("file"), viper.GetString("log"), viper.GetBool("logstash"),
			viper.GetString("batchId"), viper.GetString("cfgs"))
		if err != nil {
			return err
		}
		err = up.UploadFile(viper.GetString("url"), viper.GetString("auth"), viper.GetString("file"), viper.GetString("log"), viper.GetString("host"), viper.GetString("svcgrp"),
			viper.GetString("dtz"), viper.GetString("ids"), viper.GetString("cfgs"), viper.GetString("tags"),
			viper.GetString("batchId"), viper.GetBool("nobatch"), viper.GetBool("logstash"), version)
		if err != nil {
			return err
		}
		_, err = fmt.Fprintln(cmd.OutOrStdout(), "Upload Completed successfully")
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().StringP("file", "f", "", "File path to upload")
	upCmd.Flags().StringP("log", "l", "", "Logtype of file being uploaded.  Set to 'stream' if using STDIN.  Defaults to base name from file")
	upCmd.Flags().String("host", "", "Hostname or other identifier representing the source of the file being uploaded")
	upCmd.Flags().String("svcgrp", "default", "Defines a failure domain boundary for anomaly correlation. Learn more: https://docs.sciencelogic.com/zebrium/latest/Content/Web_Zebrium/Key_Concepts.html#service-groups")
	upCmd.Flags().String("dtz", "", "Time zone of the Logs")
	upCmd.Flags().String("ids", "", "Comma seperated list of key-value pairs of ids to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().String("cfgs", "", "Comma seperated list of key-value pairs of cfgs to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().String("tags", "", "Comma seperated list of key-value pairs of tags to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().StringP("batchId", "b", "", "Existing batch id to use")
	upCmd.Flags().Bool("logstash", false, "File is in the logstash format")
	upCmd.Flags().Bool("nobatch", false, "Disables batch processing for upload")
	err := viper.BindPFlags(upCmd.Flags())
	if err != nil {
		println(err)
		os.Exit(1)
	}

}
