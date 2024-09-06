### GRPC
- No Port, With Path: 🛑 API Validation Error
- No Port: ❌ Crash Loop
- With Port, With Path: 🛑 API Validation Error
- With Port: ✅ No Error

### HTTP
- No Port, With Path:
  - With Scheme: ✅ No Error
  - Without Scheme: ❌ Unsupported protocol scheme
- No Port:
  - With Scheme: ✅ No Error
  - Without Scheme: ❌ Unsupported protocol scheme
- With Port, With Path: 
  - With Scheme: ✅ No Error
  - Without Scheme: ❌ Unsupported protocol scheme, ❌ Wrong path join 
- With Port: 
  - With Scheme: ✅ No Error
  - Without Scheme: ❌ Unsupported protocol scheme


| Protocol | With Port | With Path | With Scheme | Result                                         |
| -------- | --------- | --------- | ----------- | ---------------------------------------------- |
| GRPC     | ❌         | ✅         | Both        | 🛑 API Validation Error                         |
| GRPC     | ❌         | ❌         | Both        | 🔺 Crash Loop                                   |
| GRPC     | ✅         | ✅         | Both        | 🛑 API Validation Error                         |
| GRPC     | ✅         | ❌         | Both        | No Error                                       |
|          |           |           |             |                                                |
| HTTP     | ❌         | ✅         | ✅           | No Error                                       |
| HTTP     | ❌         | ✅         | ❌           | 🔺 Unsupported protocol scheme                  |
| HTTP     | ❌         | ❌         | ✅           | No Error                                       |
| HTTP     | ❌         | ❌         | ❌           | 🔺 Unsupported protocol scheme                  |
| HTTP     | ✅         | ✅         | ✅           | No Error                                       |
| HTTP     | ✅         | ✅         | ❌           | 🔺 Unsupported protocol scheme, Wrong path join |
| HTTP     | ✅         | ❌         | ✅           | No Error                                       |
| HTTP     | ✅         | ❌         | ❌           | 🔺 Unsupported protocol scheme                  |