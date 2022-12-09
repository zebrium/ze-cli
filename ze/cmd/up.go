/*
Copyright © 2022 ScienceLogic Inc
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/zebrium/ze-cli/ze/common"
	"github.com/zebrium/ze-cli/ze/up"
	"log"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Upload a file to Zebrium",
	Long:  `Uploads a file to or tar file to Zebrium for analysis`,
	Run: func(cmd *cobra.Command, args []string) {
		common.ValidateAuthToken(viper.GetString("auth"))
		common.ValidateZapiUrl(viper.GetString("url"))
		common.ValidateUpMetadata(viper.GetString("filename"), viper.GetString("logtype"), viper.GetBool("logstash"),
			viper.GetString("batchId"), viper.GetString("cfgs"))
		file, _ := cmd.Flags().GetString("file")
		err := up.UploadFile(viper.GetString("url"), viper.GetString("auth"), file, viper.GetString("logtype"), viper.GetString("host"), viper.GetString("svcgrp"),
			viper.GetString("dtz"), viper.GetString("ids"), viper.GetString("cfgs"), viper.GetString("tags"),
			viper.GetString("batchId"), viper.GetBool("nobatch"), viper.GetBool("logstash"), version)
		if err != nil {
			log.Fatal(err.Error())
		}

	},
}

func init() {
	rootCmd.AddCommand(upCmd)
	upCmd.Flags().StringP("file", "f", "", "File path to upload")
	upCmd.Flags().StringP("logtype", "l", "", "Logtype of file being uploaded.  Set to 'stream' if using STDIN.  Defaults to base name from file")
	upCmd.Flags().String("host", "", "Hostname or other identifier representing the source of the file being uploaded")
	upCmd.Flags().String("svcgrp", "default", "Defines a failure domain boundary for anomaly correlation. Learn more: https://docs.zebrium.com/docs/concepts/service-group")
	upCmd.Flags().String("dtz", "", "Time zone of the Logs")
	upCmd.Flags().String("ids", "", "Comma seperated list of key-value pairs of ids to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().String("cfgs", "", "Comma seperated list of key-value pairs of cfgs to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().String("tags", "", "Comma seperated list of key-value pairs of tags to add.  eg: name1=val1,name2=val2")
	upCmd.Flags().StringP("batchId", "b", "", "Existing batch id to use")
	upCmd.Flags().Bool("logstash", false, "File is in the logstash format")
	upCmd.Flags().Bool("nobatch", false, "Disables batch processing for upload")
	viper.BindPFlags(upCmd.Flags())

}
