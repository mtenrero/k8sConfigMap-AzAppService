package main

import (
	"flag"
	"log"
)

func main() {
	var filePath = flag.String("path", "./configmap.yaml", "Specify ConfigMap file path")
	flag.Parse()
	log.Println(filePath)
}
