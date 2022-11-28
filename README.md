# loggrepper
For filtering and aggregating trace log files

## Example usages

Currently filtering logs with traceID and search strings is supported:
```
./loggrepper grep traces <traceids> --file agent-log-0 --outfile filteredTraces
```

For searching and filtering traces from the log file which contain some keywords or substrings, we can
provide a list of strings with ``--withsubstr`` flag:
```
./loggrepper grep traces --withsubstr testcluster, samplecluster --file agent-log-0 --outfile filteredTracesWithSearchString
```
There is also an option to filter traces into a Jaeger/Zipkin compatible JSON output file, which can be loaded into Jaeger UI for better visualization. You need to provide the ``--json`` or ``-j`` flag to convert it into JSON format:
```
./loggrepper convert --file agent-log-0 --json --outfile output.json
```

The ``--json`` flag can also be provided with ``loggrepper grep traces`` command.