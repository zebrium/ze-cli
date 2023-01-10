// Package cmd Copyright © 2023 ScienceLogic Inc/*
package cmd

import (
	"github.com/spf13/cobra"
)

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "A list of commands for interacting with batches",
	Long: `Zebrium batch uploads provide a way for grouping one or more related uploads so that they can be monitored and managed later as a unit. 
Each batch has a unique id used to identify the batch.`,
}

func init() {
	rootCmd.AddCommand(batchCmd)
}
