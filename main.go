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

	log.Println("Parsing ConfigMap YAML file...")

	envVars, err := cmapparser.ParseConfigMapEnvVars(*filePath)
	if err != nil {
		log.Fatal("Error parsing YAML file")
	}

	log.Println("Converting to appropriate YAML format")

	jsonEnvVars := cmapparser.ConvertToAzureJSON(envVars)
	jsonBytes, err := json.Marshal(jsonEnvVars)
	if err != nil {
		log.Fatal("Error generating JSON output file")
	}

	log.Println("Writing result to the temporary file...")
	err = ioutil.WriteFile(*outPath, jsonBytes, 0644)
	if err != nil {
		log.Println("Succeeded! :)")
	} else {
		log.Println(err)
	}
}
