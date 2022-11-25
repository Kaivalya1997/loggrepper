package structs

type Span struct {
	// the ID for the root span
	// which is also the ID for the trace itself
	TraceID string `json:"traceid"`

	// For the root span, this will be equal
	// to the TraceId
	SpanID string `json:"id"`

	// For the root span, this will be <= 0
	ParentID string `json:"parentid,omitempty"`

	Name string `json:"name"`

	Entity string `json:"entity,omitempty"`

	Duration string `json:"duration"`

	Annotations []*Annotation `json:"Annotations,omitempty"`
}

type Annotation struct {
	Time string `json:"Time"`

	Message string `json:"Message"`

	Attributes *CallerObj `json:"Attributes,omitempty"`
}

type CallerObj struct {
	CallerLocation string `json:"CallerLocation,omitempty"`
}

//JSON format structs

type TraceJson struct {
	TraceID string `json:"traceID"`

	Spans     []*SpanJson         `json:"spans"`
	Processes map[string]*Process `json:"processes"`
	Warnings  interface{}         `json:"warnings"`
}

type SpanJson struct {
	TraceID       string       `json:"traceID"`
	SpanID        string       `json:"spanID"`
	Flags         int          `json:"flags"`
	StartTime     int64        `json:"startTime,omitempty"`
	OperationName string       `json:"operationName,omitempty"`
	References    []*Reference `json:"references"`
	Duration      int64        `json:"duration,omitempty"`
	Tags          []*Tag       `json:"tags,omitempty"`
	Logs          []*Log       `json:"logs,omitempty"`
	ProcessID     string       `json:"processID"`
	Warnings      interface{}  `json:"warnings"`
}

type Reference struct {
	RefType string `json:"refType,omitempty"`
	TraceID string `json:"traceID,omitempty"`
	SpanID  string `json:"spanID,omitempty"`
}

type Log struct {
	Timestamp int64  `json:"timestamp,omitempty"`
	Fields    []*Tag `json:"fields,omitempty"`
}

type Tag struct {
	Key   string `json:"key,omitempty"`
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type JsonData struct {
	Data   []TraceJson `json:"data"`
	Total  int         `json:"total"`
	Limit  int         `json:"limit"`
	Offset int         `json:"offset"`
	Errors interface{} `json:"errors"`
}

type Process struct {
	ServiceName string `json:"serviceName"`
}
