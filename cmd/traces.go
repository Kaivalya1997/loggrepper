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
	"github.com/Kaivalya1997/loggrepper/cmd/structs"
	"github.com/spf13/cobra"
)

// tracesCmd represents the traces command
var tracesCmd = &cobra.Command{
	Use:   "traces",
	Short: "searches for traces using traceID",
	Run: func(cmd *cobra.Command, args []string) {
		inputfile, _ := cmd.Flags().GetString("file")
		outputfile, _ := cmd.Flags().GetString("outfile")
		isJson, _ := cmd.Flags().GetBool("json")
		searchStr, _ := cmd.Flags().GetStringSlice("withsubstr")

		filterTraces(inputfile, outputfile, searchStr, args, isJson)
	},
}

func init() {
	grepCmd.AddCommand(tracesCmd)
	searchStrings := []string{}
	tracesCmd.Flags().BoolP("json", "j", false, "To convert to JSON format")
	tracesCmd.Flags().StringSliceVarP(&searchStrings, "withsubstr", "s", []string{}, "Filters traces which contain substrings")
}

func filterTraces(inputFile string, outputFile string, searchStr []string, traceIds []string, isJson bool) {
	spans, spansJsonObj, traces := common.PopulateStructsFromFile(inputFile)
	idSet := make(map[string]bool)
	if len(searchStr) > 0 {
		spanSlice := []*structs.Span{}
		for _, val := range spans {
			spanSlice = append(spanSlice, val)
		}
		traceIdsFromSearch := common.FilterTracesWithSearchString(spanSlice, searchStr)
		//if searchStr is provided we ignore provided traceIds
		traceIds = traceIdsFromSearch
	}
	for _, elem := range traceIds {
		idSet[elem] = true
	}
	if !isJson {
		filteredSpans := []*structs.Span{}
		for _, span := range spans {
			if idSet[span.TraceID] {
				filteredSpans = append(filteredSpans, span)
			}
		}
		common.WriteSpansToFile(filteredSpans, outputFile)
		return
	}

	common.EstablishChildReferences(spans, spansJsonObj)
	filteredTraces := make(map[string][]*structs.SpanJson)
	for id, val := range traces {
		if idSet[id] {
			filteredTraces[id] = val
		}
	}
	common.PopulateJsonFile(filteredTraces, outputFile)
}
