### 📊 30 Apr. 2025 Runs - Total (Unfiltered)

|  Run  |     Provider     |             Scenario              | [Istio] Requests Total | [Istio] Request Duration (ms) | [Istio] Request/Response Bytes | [Istio] Request/Response Messages | [K8S] Received Bandwidth (min-max, KB/s) | [K8S] Transmitted Bandwidth (KB/s) | [K8S] Packets Rate (Received/Transmitted) | [K8S] Packets Dropped (Received + Transmitted) | [K8S] CPU Usage (istio-proxy, nginx containers) | [K8S] Memory Usage (WSS) (istio-proxy, nginx containers) |               Observed Timerange (CEST/UTC+2)               |
| :---: | :--------------: | :-------------------------------: | :--------------------: | :---------------------------: | :----------------------------: | :-------------------------------: | :--------------------------------------: | :--------------------------------: | :---------------------------------------: | :--------------------------------------------: | :---------------------------------------------: | :------------------------------------------------------: | :---------------------------------------------------------: |
|  R01  |    kyma-logs     |            Functional             |          1037          |             4.63              |          1037 / 1037           |            28.5 / 28.5            |                439 - 454                 |                1230                |               566 / 665 p/s               |                     0 p/s                      |        istio-proxy: 0.249, nginx: 0.061         |          istio-proxy: 44.8 MiB, nginx: 4.47 MiB          | `{"from":"2025-04-30 15:13:00","to":"2025-04-30 15:23:00"}` |
|  R02  | telemetry-stdout |            Functional             |          1128          |             4.08              |          1129 / 1129           |               0 / 0               |                482 - 508                 |                827                 |               566 / 649 p/s               |                     0 p/s                      |        istio-proxy: 0.250, nginx: 0.063         |          istio-proxy: 49.8 MiB, nginx: 4.48 MiB          | `{"from":"2025-04-30 15:30:00","to":"2025-04-30 15:40:00"}` |
|  R03  |    kyma-logs     |       Backend not reachable       |          956           |             4.13              |           952 / 952            |               0 / 0               |     478 - 478 (10.5 spike at 15:46)      |                719                 |               493 / 563 p/s               |                     0 p/s                      |        istio-proxy: 0.251, nginx: 0.0626        |          istio-proxy: 50.5 MiB, nginx: 4.47 MiB          | `{"from":"2025-04-30 15:45:00","to":"2025-04-30 15:55:00"}` |
|  R04  |    kyma-logs     | Backend refusing some access logs |          1073          |             4.41              |          1074 / 1074           |            29.5 / 5.15            |                449 - 467                 |                1280                |               584 / 683 p/s               |                     0 p/s                      |        istio-proxy: 0.250, nginx: 0.058         |          istio-proxy: 50.7 MiB, nginx: 4.47 MiB          | `{"from":"2025-04-30 16:08:00","to":"2025-04-30 16:18:00"}` |

> Mentions:
> - "[K8S] ..." metrics are measured on the `nginx` pod.
> - "[K8S] CPU Throttling" was always 100% for the `istio-proxy` container of the `nginx` pod.
> - The majority of the observed metrics are reported as a rate (per second) over the 10 min timerange.
> - "[Istio] Request Duration (ms)" is reported as sum/count (i.e. `istio_request_duration_milliseconds_sum` / `istio_request_duration_milliseconds_count`) over the 10 min timerange.

### 📊 30 Apr. 2025 Runs - Specific

Observed Timeranges (CEST/UTC+2):
- R01: `{"from":"2025-04-30 15:13:00","to":"2025-04-30 15:23:00"}`
- R02: `{"from":"2025-04-30 15:30:00","to":"2025-04-30 15:40:00"}`
- R03: `{"from":"2025-04-30 15:45:00","to":"2025-04-30 15:55:00"}`
- R04: `{"from":"2025-04-30 16:08:00","to":"2025-04-30 16:18:00"}`

#### `nginx` Pod
|  Run  |     Provider     |             Scenario              | [Istio] Requests Total | [Istio] Request Duration (ms) | [Istio] Request/Response Bytes | [K8S] Received/Transmitted Bandwidth (KB/s) | [K8S] Packets Rate (Received/Transmitted) | [K8S] Packets Dropped (Received + Transmitted) | [K8S] CPU Usage (istio-proxy, nginx) | [K8S] CPU Throttling (if any) | [K8S] Memory Usage (WSS) (istio-proxy, nginx) |
| :---: | :--------------: | :-------------------------------: | :--------------------: | :---------------------------: | :----------------------------: | :-----------------------------------------: | :---------------------------------------: | :--------------------------------------------: | :----------------------------------: | :---------------------------: | :-------------------------------------------: |
|  R01  |    kyma-logs     |            Functional             |          504           |             2.36              |           504 / 504            |                 439 / 1230                  |               566 / 665 p/s               |                     0 p/s                      |   istio-proxy: 0.249, nginx: 0.061   |       istio-proxy: 100%       |    istio-proxy: 44.8 MiB, nginx: 4.47 MiB     |
|  R02  | telemetry-stdout |            Functional             |          564           |             1.95              |           564 / 564            |                  485 / 827                  |               566 / 649 p/s               |                     0 p/s                      |   istio-proxy: 0.250, nginx: 0.063   |       istio-proxy: 100%       |    istio-proxy: 49.8 MiB, nginx: 4.48 MiB     |
|  R03  |    kyma-logs     |       Backend not reachable       |          490           |             2.01              |           486 / 486            |                  478 / 719                  |               493 / 563 p/s               |                     0 p/s                      |  istio-proxy: 0.251, nginx: 0.0626   |       istio-proxy: 100%       |    istio-proxy: 50.5 MiB, nginx: 4.47 MiB     |
|  R04  |    kyma-logs     | Backend refusing some access logs |          522           |              2.2              |           522 / 522            |                 463 / 1280                  |               584 / 683 p/s               |                     0 p/s                      |   istio-proxy: 0.250, nginx: 0.058   |       istio-proxy: 100%       |    istio-proxy: 50.7 MiB, nginx: 4.47 MiB     |

#### `fortio` Pod
|  Run  |     Provider     |             Scenario              | [Istio] Requests Total | [Istio] Request Duration (ms) | [Istio] Request/Response Bytes | [K8S] Received/Transmitted Bandwidth (KB/s) | [K8S] Packets Rate (Received/Transmitted) | [K8S] Packets Dropped (Received + Transmitted) | [K8S] CPU Usage (istio-proxy, fortio) | [K8S] CPU Throttling (if any) | [K8S] Memory Usage (WSS) (istio-proxy, fortio) |
| :---: | :--------------: | :-------------------------------: | :--------------------: | :---------------------------: | :----------------------------: | :-----------------------------------------: | :---------------------------------------: | :--------------------------------------------: | :-----------------------------------: | :---------------------------: | :--------------------------------------------: |
|  R01  |    kyma-logs     |            Functional             |          504           |             7.04              |           504 / 504            |                  728 / 446                  |               589 / 509 p/s               |                     0 p/s                      |   istio-proxy: 0.17, fortio: 0.0497   |        istio-proxy: 0%        |    istio-proxy: 39.5 MiB, fortio: 10.6 MiB     |
|  R02  | telemetry-stdout |            Functional             |          564           |             6.21              |           564 / 564            |                  795 / 496                  |               646 / 566 p/s               |                     0 p/s                      |  istio-proxy: 0.172, fortio: 0.0538   |        istio-proxy: 0%        |    istio-proxy: 40.6 MiB, fortio: 11.0 MiB     |
|  R03  |    kyma-logs     |       Backend not reachable       |          466           |             6.34              |           466 / 466            |                  818 / 421                  |               548 / 480 p/s               |                     0 p/s                      |   istio-proxy: 0.17, fortio: 0.0528   |        istio-proxy: 0%        |    istio-proxy: 40.5 MiB, fortio: 10.7 MiB     |
|  R04  |    kyma-logs     | Backend refusing some access logs |          522           |              6.8              |           522 / 522            |                  748 / 459                  |               605 / 525 p/s               |                     0 p/s                      |   istio-proxy: 0.17, fortio: 0.0509   |        istio-proxy: 0%        |    istio-proxy: 40.4 MiB, fortio: 11.4 MiB     |



nginx: 100.64.0.27

fortio:
- R01: 100.64.0.28
- R02: 100.64.0.29
- R03: 100.64.0.30
- R04: 100.64.0.33