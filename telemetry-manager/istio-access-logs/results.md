## Observed metrics
Envoy: https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/statistics
Istio: https://istio.io/latest/docs/reference/config/metrics/

- `envoy_server_memory_allocated` = Current amount of allocated memory in bytes. Total of both new and old Envoy processes on hot restart.
- `envoy_server_memory_heap_size` = Current reserved heap size in bytes. New Envoy process heap size on hot restart.
- `envoy_server_memory_physical_size` = Current estimate of total bytes of the physical memory. New Envoy process physical memory size on hot restart.
- `envoy_server_total_connections` = Total connections of both new and old Envoy processes.
- (TBD) `envoy_server_initialization_time_ms` = Total time taken for Envoy initialization in milliseconds. This is the time from server start-up until the worker threads are ready to accept new connections.
- `envoy_server_envoy_bug_failures` = Number of envoy bug failures detected in a release build. File or report the issue if this increments as this may be serious.
- load-test/nginx/istio-proxy: CPU Usage, CPU Throttling, Memory Usage (`Dashboards > Kubernetes / Compute Resources / Pod`)

Istio:
- Request Count (`istio_requests_total`): This is a COUNTER incremented for every request handled by an Istio proxy.
- Request Duration (`istio_request_duration_milliseconds`): This is a DISTRIBUTION which measures the duration of requests.
- Request Size (`istio_request_bytes`): This is a DISTRIBUTION which measures HTTP request body sizes.
- Response Size (`istio_response_bytes`): This is a DISTRIBUTION which measures HTTP response body sizes.
- gRPC Request Message Count (`istio_request_messages_total`): This is a COUNTER incremented for every gRPC message sent from a client.
- gRPC Response Message Count (`istio_response_messages_total`): This is a COUNTER incremented for every gRPC message sent from a server.

## Results

### Run R001
- 22-Apr-2025, 13:50-14:10
- Load Generator (fortio): `["load", "-t", "10m", "-qps", "0", "-nocatchup", "-uniform", "nginx.load-test.svc"]`
#### kyma-logs
- `envoy_server_memory_allocated` / `envoy_server_memory_heap_size` / `envoy_server_memory_physical_size` (MB): 10.1 / 16 / 21.3
- `envoy_server_total_connections` (average): 3.25
- `envoy_server_envoy_bug_failures`: 0
- `envoy_server_uptime`: 8.4K - 9.4K
- `envoy_cluster_upstream_cx_max_requests`: 7.25
- `envoy_cluster_upstream_cx_protocol_error`: 6.25
- `envoy_cluster_upstream_cx_destroy`: 6.25
- ![envoy_cluster_upstream_bytes](results/R001_envoy_cluster_upstream_bytes_kyma-logs.png)
| Pod         | CPU Usage | CPU Throttling | Memory Usage (WSS) | Receive Bandwidth | Transmit Bandwidth | Rate of Received Packets | Rate of Transmitted Packets | Rate of Packets Dropped (Received + Transmitted) |
| ----------- | --------- | -------------- | ------------------ | ----------------- | ------------------ | ------------------------ | --------------------------- | ------------------------------------------------ |
| istio-proxy | 0.25      | 100%           | 44.0 MiB           | -                 | -                  | -                        | -                           | -                                                |
| nginx       | ~0.06     | -              | 4.46 MiB           | 467 KB/s          | 1.3 MB/s           | 595 p/s                  | 690 p/s                     | 0 p/s                                            |

fortio logs:
```
2025-04-22T12:00:31.826092066Z fortio {"ts":1745323231.825889,"level":"info","r":1,"file":"scli.go","line":122,"msg":"Starting","command":"Φορτίο","version":"1.69.4 h1:G0DXdTn8/QtiCh+ykBXft8NcOCojfAhQKseHuxFVePE= go1.23.8 amd64 linux","go-max-procs":4}
2025-04-22T12:00:31.826722459Z fortio Fortio 1.69.4 running at 0 queries per second, 4->4 procs, for 10m0s: nginx.load-test.svc
2025-04-22T12:00:31.826731356Z fortio {"ts":1745323231.826656,"level":"info","r":1,"file":"httprunner.go","line":121,"msg":"Starting http test","run":0,"url":"nginx.load-test.svc","threads":4,"qps":"-1.0","warmup":"parallel","conn-reuse":""}
2025-04-22T12:00:31.826852744Z fortio {"ts":1745323231.826762,"level":"warn","r":1,"file":"http_client.go","line":172,"msg":"Assuming http:// on missing scheme for 'nginx.load-test.svc'"}
2025-04-22T12:00:31.838827230Z fortio Starting at max qps with 4 thread(s) [gomax 4] for 10m0s
...
2025-04-22T12:10:31.863270576Z fortio {"ts":1745323831.863032,"level":"info","r":23,"file":"periodic.go","line":851,"msg":"T003 ended after 10m0.024321553s : 79617 calls. qps=132.68962130390483"}
2025-04-22T12:10:31.863313822Z fortio {"ts":1745323831.863077,"level":"info","r":22,"file":"periodic.go","line":851,"msg":"T002 ended after 10m0.02437356s : 79600 calls. qps=132.6612776206504"}
2025-04-22T12:10:31.863320783Z fortio {"ts":1745323831.863032,"level":"info","r":21,"file":"periodic.go","line":851,"msg":"T001 ended after 10m0.02432148s : 79612 calls. qps=132.6812883244994"}
2025-04-22T12:10:31.863326224Z fortio {"ts":1745323831.863191,"level":"info","r":20,"file":"periodic.go","line":851,"msg":"T000 ended after 10m0.024501371s : 79679 calls. qps=132.79291065271653"}
2025-04-22T12:10:31.863351196Z fortio Ended after 10m0.024541595s : 318508 calls. qps=530.82
2025-04-22T12:10:31.863532733Z fortio {"ts":1745323831.863323,"level":"info","r":1,"file":"periodic.go","line":581,"msg":"Run ended","run":0,"elapsed":600024541595,"calls":318508,"qps":530.8249545149174}
2025-04-22T12:10:31.863553727Z fortio Aggregated Function Time : count 318508 avg 0.0075341122 +/- 0.01669 min 0.001008034 max 0.0788881 sum 2399.67501
2025-04-22T12:10:31.863560390Z fortio # range, mid point, percentile, count
2025-04-22T12:10:31.863591553Z fortio >= 0.00100803 <= 0.002 , 0.00150402 , 8.22, 26167
2025-04-22T12:10:31.863598347Z fortio > 0.002 <= 0.003 , 0.0025 , 66.35, 185163
2025-04-22T12:10:31.863602940Z fortio > 0.003 <= 0.004 , 0.0035 , 87.90, 68639
2025-04-22T12:10:31.863607931Z fortio > 0.004 <= 0.005 , 0.0045 , 91.06, 10059
2025-04-22T12:10:31.863613043Z fortio > 0.005 <= 0.006 , 0.0055 , 91.87, 2575
2025-04-22T12:10:31.863617678Z fortio > 0.006 <= 0.007 , 0.0065 , 92.19, 1040
2025-04-22T12:10:31.863622400Z fortio > 0.007 <= 0.008 , 0.0075 , 92.33, 437
2025-04-22T12:10:31.863626960Z fortio > 0.008 <= 0.009 , 0.0085 , 92.40, 213
2025-04-22T12:10:31.863631754Z fortio > 0.009 <= 0.01 , 0.0095 , 92.43, 96
2025-04-22T12:10:31.863636137Z fortio > 0.01 <= 0.011 , 0.0105 , 92.44, 28
2025-04-22T12:10:31.863640978Z fortio > 0.011 <= 0.012 , 0.0115 , 92.45, 31
2025-04-22T12:10:31.863645733Z fortio > 0.012 <= 0.014 , 0.013 , 92.45, 19
2025-04-22T12:10:31.863763274Z fortio > 0.014 <= 0.016 , 0.015 , 92.46, 20
2025-04-22T12:10:31.863885662Z fortio > 0.016 <= 0.018 , 0.017 , 92.46, 5
2025-04-22T12:10:31.863900708Z fortio > 0.018 <= 0.02 , 0.019 , 92.46, 3
2025-04-22T12:10:31.863906658Z fortio > 0.02 <= 0.025 , 0.0225 , 92.46, 2
2025-04-22T12:10:31.863911576Z fortio > 0.025 <= 0.03 , 0.0275 , 92.47, 15
2025-04-22T12:10:31.863916704Z fortio > 0.03 <= 0.035 , 0.0325 , 92.47, 10
2025-04-22T12:10:31.863921410Z fortio > 0.035 <= 0.04 , 0.0375 , 92.47, 14
2025-04-22T12:10:31.863926040Z fortio > 0.04 <= 0.045 , 0.0425 , 92.48, 18
2025-04-22T12:10:31.863938315Z fortio > 0.045 <= 0.05 , 0.0475 , 92.49, 47
2025-04-22T12:10:31.863943492Z fortio > 0.05 <= 0.06 , 0.055 , 93.10, 1924
2025-04-22T12:10:31.863961030Z fortio > 0.06 <= 0.07 , 0.065 , 98.66, 17727
2025-04-22T12:10:31.863966467Z fortio > 0.07 <= 0.0788881 , 0.0744441 , 100.00, 4256
2025-04-22T12:10:31.863993811Z fortio # target 50% 0.00271876
2025-04-22T12:10:31.863998422Z fortio # target 75% 0.00340139
2025-04-22T12:10:31.864002274Z fortio # target 90% 0.0046649
2025-04-22T12:10:31.864006049Z fortio # target 99% 0.0722365
2025-04-22T12:10:31.864009653Z fortio # target 99.9% 0.0782229
2025-04-22T12:10:31.864013167Z fortio Error cases : no data
2025-04-22T12:10:31.864033292Z fortio # Socket and IP used for each connection:
2025-04-22T12:10:31.864039149Z fortio [0]   1 socket used, resolved to 100.106.55.77:80, connection timing : count 1 avg 0.000311963 +/- 0 min 0.000311963 max 0.000311963 sum 0.000311963
2025-04-22T12:10:31.864043269Z fortio [1]   1 socket used, resolved to 100.106.55.77:80, connection timing : count 1 avg 0.000188963 +/- 0 min 0.000188963 max 0.000188963 sum 0.000188963
2025-04-22T12:10:31.864047394Z fortio [2]   1 socket used, resolved to 100.106.55.77:80, connection timing : count 1 avg 9.9328e-05 +/- 0 min 9.9328e-05 max 9.9328e-05 sum 9.9328e-05
2025-04-22T12:10:31.864051310Z fortio [3]   1 socket used, resolved to 100.106.55.77:80, connection timing : count 1 avg 0.000465303 +/- 0 min 0.000465303 max 0.000465303 sum 0.000465303
2025-04-22T12:10:31.864055751Z fortio Connection time histogram (s) : count 4 avg 0.00026638925 +/- 0.0001374 min 9.9328e-05 max 0.000465303 sum 0.001065557
2025-04-22T12:10:31.864060293Z fortio # range, mid point, percentile, count
2025-04-22T12:10:31.864064692Z fortio >= 9.9328e-05 <= 0.000465303 , 0.000282316 , 100.00, 4
2025-04-22T12:10:31.864069702Z fortio # target 50% 0.00022132
2025-04-22T12:10:31.864073527Z fortio # target 75% 0.000343311
2025-04-22T12:10:31.864086867Z fortio # target 90% 0.000416506
2025-04-22T12:10:31.864091216Z fortio # target 99% 0.000460423
2025-04-22T12:10:31.864147473Z fortio # target 99.9% 0.000464815
2025-04-22T12:10:31.864154004Z fortio Sockets used: 4 (for perfect keepalive, would be 4)
2025-04-22T12:10:31.864158370Z fortio Uniform: true, Jitter: false, Catchup allowed: false
2025-04-22T12:10:31.864162769Z fortio IP addresses distribution:
2025-04-22T12:10:31.864166821Z fortio 100.106.55.77:80: 4
2025-04-22T12:10:31.864170706Z fortio Code 200 : 318508 (100.0 %)
2025-04-22T12:10:31.864174991Z fortio Response Header Sizes : count 318508 avg 241.07537 +/- 0.264 min 241 max 242 sum 76784434
2025-04-22T12:10:31.864178830Z fortio Response Body/Total Sizes : count 318508 avg 856.07537 +/- 0.264 min 856 max 857 sum 272666854
2025-04-22T12:10:31.864182392Z fortio All done 318508 calls (plus 4 warmup) 7.534 ms avg, 530.8 qps
2025-04-22T12:10:32.10588Z     stream closed EOF for load-test/traffic-generator (fortio)
```
#### envoy