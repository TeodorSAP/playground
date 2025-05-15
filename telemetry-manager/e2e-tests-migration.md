# Benchmarking Migrated E2E Tests

## Overview
| Batch Name | Runs | Avg. Duration                        |
| ---------- | ---- | ------------------------------------ |
| fluentbit  | ✅✅✅  | (144.63 + 63.86 + 93.8) / 3 = 100.76 |
| agent      | ✅??  | ?                                    |
| gateway    | ???  | ?                                    |
| shared     | ???  | ?                                    |
| ALL        | ???  | ?                                    |

### Command
`go test -timeout 10m -count 1 github.com/kyma-project/telemetry-manager/test/e2e/logs/migrated/fluentbit`
> - `-count 1`: Disable test caching
> - `-timeout 10m`: Set a timeout of 10 minutes for the test run (default is 10 minutes)

## Verbose Output

### fluentbit
```shell
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

### agent
```shell
=== RUN   TestAttributesParser
--- PASS: TestAttributesParser (28.19s)
PASS
ok      github.com/kyma-project/telemetry-manager/test/e2e/logs/migrated/agent  28.636s
```

### gateway
```shell
```

### shared
```shell
```

### ALL
```shell