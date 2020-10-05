# k8sconfigmap-azappservice

[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
![Latest CI Status](https://badgen.net/github/status/mtenrero/k8sconfigmap-azappservice/main)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gomods/athens.svg)](https://github.com/gomods/athens)
[![Coverage Status](https://coveralls.io/repos/github/mtenrero/k8sConfigMap-AzAppService/badge.svg?branch=main)](https://coveralls.io/github/mtenrero/k8sConfigMap-AzAppService?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtenrero/k8sconfigmap-azappservice)](https://goreportcard.com/report/github.com/mtenrero/k8sconfigmap-azappservice)
[![Latest Release](https://badgen.net/github/release/mtenrero/k8sconfigmap-azappservice/stable)](https://badgen.net/github/release/mtenrero/k8sconfigmap-azappservice/stable)

This CLI executable program, transform a given Kubernetes ConfigMap into a JSON object ready to consume by Azure
AppService container as environment variables.

It's quite useful for avoiding to maintain multiple config files if you have hybrid deployment with kubernetes and 
Azure App Service containers.

## Usage

Download the latest binary for your platform, (its written in go and doesn't need any dependency, it's all compiled!)

Execute it in your CI/CD pipeline:

```bash
./k8sconfigmap-azappservice -path <k8s-configmap-file> -out <json-output-file>
```

## Contributing 

PR! are accepted. Its planned to improve the program for supporting multiple ConfigMaps in the same file and the ability
to specify environment-scoped variables for multiple apps deployment inside the same Azure App Service.
