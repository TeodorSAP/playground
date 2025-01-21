| Test                                          | Result | Duration | CPU Avg% | CPU Max% | RAM Avg MiB | RAM Max MiB | Sent Items | Received Items | Log Records / sec |
| --------------------------------------------- | ------ | -------: | -------: | -------: | ----------: | ----------: | ---------: | -------------: | ----------------: |
| IdleMode                                      | PASS   |      17s |      0.6 |      2.3 |          46 |          65 |          0 |              0 |                   |
| Log10kDPS/OTLP                                | PASS   |      16s |     14.0 |     18.1 |          59 |          85 |     150100 |         150100 |                   |
| Log10kDPS/OTLP-HTTP                           | PASS   |      15s |      9.1 |     63.3 |          56 |          77 |     150100 |         150100 |                   |
| Log10kDPS/filelog                             | PASS   |      15s |     12.1 |     13.5 |          60 |          86 |     150100 |         150100 |                   |
| Log10kDPS/filelog_checkpoints                 | PASS   |      15s |     11.3 |     11.6 |          60 |          85 |     150100 |         150100 |                   |
| Log10kDPS/kubernetes_containers               | PASS   |      15s |     23.2 |     25.3 |          66 |          95 |     150100 |         150100 |                   |
| Log10kDPS/kubernetes_containers_parser        | PASS   |      15s |     23.1 |     72.8 |          60 |          85 |     150000 |         150000 |                   |
| Log10kDPS/k8s_CRI-Containerd                  | PASS   |      15s |     21.5 |     22.8 |          62 |          89 |     150000 |         150000 |                   |
| Log10kDPS/k8s_CRI-Containerd_no_attr_ops      | PASS   |      15s |     20.3 |     64.4 |          66 |          95 |     150100 |         150100 |                   |
| Log10kDPS/CRI-Containerd                      | PASS   |      15s |     13.1 |     13.9 |          64 |          91 |     150100 |         150100 |                   |
| Log10kDPS/syslog-tcp-batch-1                  | PASS   |      16s |     22.9 |     43.6 |          57 |          82 |     150100 |         150100 |                   |
| Log10kDPS/syslog-tcp-batch-100                | PASS   |      15s |      8.1 |     10.0 |          57 |          82 |     150100 |         150100 |                   |
| Log10kDPS/FluentForward-SplunkHEC             | PASS   |      16s |     31.9 |     69.9 |          65 |          89 |     150100 |         150100 |             9.381 |
| Log10kDPS/tcp-batch-1                         | PASS   |      16s |     18.6 |     21.5 |          57 |          82 |     150000 |         150000 |             9.375 |
| Log10kDPS/tcp-batch-100                       | PASS   |      15s |      8.1 |     90.7 |          59 |          82 |     150100 |         150100 |            10.006 |
| LogLargeFiles/filelog-largefiles-2Gb-lifetime | PASS   |     100s |      0.0 |      0.0 |           0 |           0 |   18804584 |       18804584 |           188.045 |
| LogLargeFiles/filelog-largefiles-6GB-lifetime | PASS   |     201s |      0.0 |      0.0 |           0 |           0 |   57050440 |       57050440 |           283.833 |