package main

import (
	"encoding/json"
	"flag"
	"github.com/mtenrero/k8sconfigmap-azappservice/cmapparser"
	"io/ioutil"
	"log"
)

func main() {
	var filePath = flag.String("path", "./configmap.yaml", "Specify ConfigMap file path")
	var outPath = flag.String("out", "./azout.json", "Specify target JSON file path")
	flag.Parse()

	envVars, err := cmapparser.ParseConfigMapEnvVars(*filePath)
	if err != nil {
		log.Fatal("Error parsing YAML file")
	}

	jsonEnvVars := cmapparser.ConvertToAzureJSON(envVars)
	jsonBytes, err := json.Marshal(jsonEnvVars)
	if err != nil {
		log.Fatal("Eror generating JSON output file")
	}
	ioutil.WriteFile(*outPath, jsonBytes, 0644)
}
