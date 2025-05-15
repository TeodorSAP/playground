# Benchmarking Migrated E2E Tests

## Overview
| Batch Label | Avg. Duration (3 successful runs) (in seconds) |
| ----------- | ---------------------------------------------- |
| fluent-bit  | (144.63 + 63.86 + 93.8) / 3 = 100.76           |
| log-agent   | (? + ? + ?)/3 =                                |
| log-gateway | (? + ? + ?)/3 =                                |


### Command
`go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=A,B,C`
> - `-count 1`: Disable test caching
> - `-timeout 0`: Disable timeout (default is 10 minutes)

## Verbose Output

### fluent-bit
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=fluent-bit

=== RUN   TestCustomFilterDenied
--- PASS: TestCustomFilterDenied (10.00s)
=== RUN   TestCustomFilterAllowed
--- PASS: TestCustomFilterAllowed (73.19s)
=== RUN   TestCustomOutput
--- PASS: TestCustomOutput (9.94s)
=== RUN   TestDedot
--- PASS: TestDedot (11.94s)
=== RUN   TestKeepAnnotations
--- PASS: TestKeepAnnotations (31.54s)
=== RUN   TestLogParser
--- PASS: TestLogParser (4.71s)
=== RUN   TestModifyTimestampDateFormat
--- PASS: TestModifyTimestampDateFormat (2.96s)
PASS
ok      github.com/kyma-project/telemetry-manager/test/e2e/logs/migrated/fluentbit      144.634s
```

### log-agent
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=log-agent
```

See [./log-agent.out](./log-agent-1.out)
- ⚠️ `TestAttributesParser` is flaky => truncated/misformated JSON in HTTP body (see `assert.shared` and `assert.log/log_otel_matchers.go/HaveFlatOTelLogs`)

### log-gateway
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=log-gateway
```

See [./log-gateway.out](./log-gateway-1.out)