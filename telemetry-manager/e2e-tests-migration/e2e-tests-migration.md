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

## Test Sets

### fluent-bit
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=fluent-bit
```

### log-agent
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=log-agent
```

### log-gateway
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=log-gateway
```

### Max Pipeline Tests
```shell
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=max-pipeline
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=max-pipeline-agent
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=max-pipeline-gateway
go test -v -timeout 0 -count 1 ./test/e2e/logs/migrated... -- -labels=max-pipeline-fluent-bit
```