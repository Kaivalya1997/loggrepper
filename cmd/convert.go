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
	"github.com/Kaivalya1997/loggrepper/cmd/common"
	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert log file to other formats like json etc.",
	Run: func(cmd *cobra.Command, args []string) {
		inputfile, _ := cmd.Flags().GetString("file")
		outputfile, _ := cmd.Flags().GetString("outfile")
		isJson, _ := cmd.Flags().GetBool("json")
		if isJson {
			convertToJson(inputfile, outputfile)
		}
	},
}

func init() {
	rootCmd.AddCommand(convertCmd)

	// Here you will define your flags and configuration settings.

	convertCmd.PersistentFlags().String("file", "", "path to log file which needs to be converted")
	convertCmd.PersistentFlags().String("outfile", "", "path to converted file")
	convertCmd.MarkPersistentFlagRequired("file")
	convertCmd.MarkPersistentFlagRequired("outfile")

	convertCmd.Flags().BoolP("json", "j", false, "To convert to JSON format")
}

func convertToJson(inputFile string, outputFile string) {
	spans, spansJsonObj, traces := common.PopulateStructsFromFile(inputFile)
	common.PopulateChildSpans(spans, spansJsonObj)
	common.PopulateJsonFile(traces, outputFile)
}
