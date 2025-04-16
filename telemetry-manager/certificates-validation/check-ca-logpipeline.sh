#!/bin/bash

kubectl get logpipelines -A -ojsonpath="{ range .items[*] }{ @.spec.output.http.tls }{ '\n' }{ end }" | grep '"ca"' | wc -l