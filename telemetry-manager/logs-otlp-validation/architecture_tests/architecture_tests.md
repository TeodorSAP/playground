## Architecture Tests

Testing different architectures in the same load test setup to see which one is the most performant one.

## Setup
- 2 nodes
- 60 generators with 10m CPU limit
- 5 min query spans

### Scenario 1: ✅ Sending Queue, ✅ Batch Processor
![21.01.2025 15:35](image.png)

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE | AGENT RESOURCES      | BACKEND RECEIVED |
| :------------- | :------------- | :---------- | :------------------- | :--------------- |
| 27K            | 7.31K          | 2K          | M:258/255, C:1.4/1.1 | 8.38K            |

- Agent receives more data than it exports
  - Hypothesis 1: Batching stacks the metric up, not showing the correct values
  - Hypothesis 2: Some logs are lost
- Sending queue always full (2K)

### Scenario 2: ✅ Sending Queue, ✅ Batcher Exporter
![21.01.2025 16:08](image-1.png)

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE | AGENT RESOURCES      | BACKEND RECEIVED |
| :------------- | :------------- | :---------- | :------------------- | :--------------- |
| 8.27K          | 8.27K          | 1.9K        | M:122/115, C:0.5/0.5 | 8.26K            |

- Agent receives the same amount of data that it exports
- Sending queue almost full, but less than in scenario 1

### Scenario 3: ✅ Sending Queue, ❌ No Batching
![21.01.2025 16:22](image-2.png)

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE              | AGENT RESOURCES      | BACKEND RECEIVED |
| :------------- | :------------- | :----------------------- | :------------------- | :--------------- |
| 6.62K          | 6.61K          | 1.99K (max. 2.95K spike) | M:114/118, C:0.4/0.4 | 6.61K            |

- Disabling batching just decreases throughput

### Scenario 4: ❌ No Queue, ❌ No Batching
TODO: Paste image here

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE | AGENT RESOURCES | BACKEND RECEIVED |
| :------------- | :------------- | :---------- | :-------------- | :--------------- |
| ?K             | ?K             | ?K          | M:?/?, C:?/?    | ?K               |

### Scenario 5: ❌ No Queue, ✅ Batch Processor
TODO: Paste image here

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE | AGENT RESOURCES | BACKEND RECEIVED |
| :------------- | :------------- | :---------- | :-------------- | :--------------- |
| ?K             | ?K             | ?K          | M:?/?, C:?/?    | ?K               |

### Scenario 6: ❌ No Queue, ✅ Batcher Exporter
TODO: Paste image here

| AGENT RECEIVED | AGENT EXPORTED | AGENT QUEUE | AGENT RESOURCES | BACKEND RECEIVED |
| :------------- | :------------- | :---------- | :-------------- | :--------------- |
| ?K             | ?K             | ?K          | M:?/?, C:?/?    | ?K               |