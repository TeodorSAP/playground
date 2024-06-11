#!/bin/bash

kubectl get metricpipelines -A -ojsonpath="{ range .items[*] }{ @.spec.output.otlp.tls }{ '\n' }{ end }" | grep '"ca"' | wc -l