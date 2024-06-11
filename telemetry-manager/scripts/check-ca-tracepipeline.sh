#!/bin/bash

kubectl get tracepipelines -A -ojsonpath="{ range .items[*] }{ @.spec.output.otlp.tls }{ '\n' }{ end }" | grep '"ca"' | wc -l