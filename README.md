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

## Using with Jaeger UI

To install and start Jaeger on linux/WSL run:
```
docker run -d --name jaeger   -e COLLECTOR_ZIPKIN_HOST_PORT=:9411   -e COLLECTOR_OTLP_ENABLED=true   -p 6831:6831/udp   -p 6832:6832/udp   -p 5778:5778   -p 16686:16686   -p 4317:4317   -p 4318:4318   -p 14250:14250   -p 14268:14268   -p 14269:14269   -p 9411:9411   jaegert
racing/all-in-one:1.39
```
The Jaeger UI will run on http://localhost:16686

There will be an option to upload a JSON file in the UI where can can upload the above generated JSON format files containing one or more traces for visualization.