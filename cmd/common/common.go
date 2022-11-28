package common

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Kaivalya1997/loggrepper/cmd/structs"
	"github.com/Kaivalya1997/loggrepper/cmd/utils"
)

func PopulateStructsFromFile(inputFile string) (map[string]*structs.Span, map[string]*structs.SpanJson, map[string][]*structs.SpanJson) {
	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("failed to open input file with error: %+v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	const maxCapacity = 512 * 1024
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	traces := make(map[string][]*structs.SpanJson)
	spans := make(map[string]*structs.Span)
	spansJsonObj := make(map[string]*structs.SpanJson)
	for scanner.Scan() {
		line := scanner.Text()
		spanObj := &structs.Span{}
		json.Unmarshal([]byte(line), &spanObj)

		//add object to spans map
		spans[spanObj.SpanID] = spanObj
		//process the struct here
		//map span obj to spanJson obj
		spanJsonObj := &structs.SpanJson{
			TraceID:       spanObj.TraceID,
			SpanID:        spanObj.SpanID,
			OperationName: spanObj.Name,
		}
		duration, err := strconv.ParseFloat(strings.ReplaceAll(spanObj.Duration, "s", ""), 64)
		if err != nil {
			panic(err)
		}
		spanJsonObj.Duration = int64(duration * 1e6)
		tags := []*structs.Tag{
			&structs.Tag{Key: "name", Type: "string", Value: spanObj.Name},
			&structs.Tag{Key: "entity", Type: "string", Value: spanObj.Entity},
		}
		spanJsonObj.Tags = tags
		logs := []*structs.Log{}

		var startTime int64 = 1
		var offsetlogTime int64
		for i, elem := range spanObj.Annotations {
			if i == 0 {
				offsetlogTime = utils.ConvertTimeStampToUnixMicro(elem.Time)
			}
			fields := []*structs.Tag{
				&structs.Tag{Key: "Message", Type: "string", Value: elem.Message},
				&structs.Tag{Key: "CallerLocation", Type: "string", Value: elem.Attributes.CallerLocation},
			}
			logs = append(logs, &structs.Log{Timestamp: startTime + (utils.ConvertTimeStampToUnixMicro(elem.Time) - offsetlogTime), Fields: fields})
		}
		if len(logs) > 0 {
			spanJsonObj.Logs = logs
		}
		spanJsonObj.StartTime = startTime
		//populate the trace with the span objects, except the Reference attribute
		traces[spanObj.TraceID] = append(traces[spanObj.TraceID], spanJsonObj)

		//add span json object to spansJsonObj map
		spansJsonObj[spanJsonObj.SpanID] = spanJsonObj

	}
	return spans, spansJsonObj, traces
}

func EstablishChildReferences(spans map[string]*structs.Span, spansJsonObj map[string]*structs.SpanJson) {

	for _, span := range spans {
		if _, ok := spans[span.ParentID]; !ok {
			// fmt.Printf("No parent span found in the file for span: %+v\n", span)
			continue
		}
		if span.ParentID == "0000000000000000" || span.ParentID == "" {
			continue
		}
		//check for parent span and populate the ids
		ref := &structs.Reference{RefType: "CHILD_OF", TraceID: spans[span.ParentID].TraceID, SpanID: spans[span.ParentID].SpanID}
		spansJsonObj[span.SpanID].References = append(spansJsonObj[span.SpanID].References, ref)
	}
}

func PopulateJsonFile(traces map[string][]*structs.SpanJson, outputFile string) {
	file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var jsonData structs.JsonData

	for traceid, spans := range traces {
		traceObj := structs.TraceJson{TraceID: traceid, Spans: spans}
		for _, elem := range spans {
			elem.ProcessID = traceid
		}
		processObj := map[string]*structs.Process{}
		processObj[traceid] = &structs.Process{ServiceName: traceid[0:10]}
		traceObj.Processes = processObj
		jsonData.Data = append(jsonData.Data, traceObj)

	}

	//write to file
	result, err := json.Marshal(jsonData)
	if err != nil {
		fmt.Println(err)
	}
	n, err := io.WriteString(file, string(result))
	if err != nil {
		fmt.Println(n, err)
	}
}

func WriteSpansToFile(spans []*structs.Span, outputFile string) {
	file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	datawriter := bufio.NewWriter(file)
	defer datawriter.Flush()
	for _, span := range spans {
		result, err := json.Marshal(span)
		if err != nil {
			fmt.Println(err)
		}
		n, err := datawriter.WriteString(string(result) + "\n")
		if err != nil {
			fmt.Println(n, err)
		}
	}
}

func FilterTracesWithSearchString(spans []*structs.Span, searchStr []string) []string {
	traceIds := []string{}
	for _, span := range spans {
		str, _ := json.Marshal(span)
		flag := false
		for _, substr := range searchStr {
			if strings.Contains(string(str), substr) {
				flag = true
				break
			}
		}
		if flag {
			traceIds = append(traceIds, span.TraceID)
		}

	}
	return traceIds
}
