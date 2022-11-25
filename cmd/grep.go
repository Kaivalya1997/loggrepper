/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// grepCmd represents the grep command
var grepCmd = &cobra.Command{
	Use:   "grep",
	Short: "Searches for logs in a provided file and aggregates them",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grep called")
	},
}

func init() {
	rootCmd.AddCommand(grepCmd)

	grepCmd.PersistentFlags().String("file", "", "Path of the log file")
	grepCmd.PersistentFlags().String("outfile", "", "Path of the output file which will have logs aggregated")
	grepCmd.MarkPersistentFlagRequired("file")
	grepCmd.MarkPersistentFlagRequired("outfile")
}
