package cmapparser

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

// ParseConfigMap loads a yaml k8s file containing a ConfigMap and retrieve its envVars
func ParseConfigMapEnvVars(filepath string) (map[string]interface{}, error) {
	var configFile, err = getFile(filepath)
	if err != nil {
		log.Fatal("Error parsing YAML file")
	}

	err = ensureConfigMap(configFile)
	if err != nil {
		log.Fatal("The given field doesn't contain a kubernetes ConfigMap!")
	}

	envs := getEnvironments(configFile)

	return envs, err
}

func ensureConfigMap(object map[string]interface{}) error {
	if strings.Compare(object["kind"].(string), "ConfigMap") == 0 {
		return nil
	}
	return errors.New("the given YAML file there isn't a ConfigMap")
}

func getEnvironments(object map[string]interface{}) map[string]interface{} {
	return castMap(object["data"].(map[interface{}]interface{}))
}

func castMap(object map[interface{}]interface{}) map[string]interface{} {
	objectStringKeys := make(map[string]interface{})

	for key, value := range object {
		objectStringKeys[key.(string)] = value
	}

	return objectStringKeys
}

func ConvertToAzureJSON(envVars map[string]interface{}) []interface{} {
	var jsonEnvs []interface{}
	for envKey, envValue := range envVars {
		var env = map[string]interface{}{
			"name":        envKey,
			"value":       envValue,
			"slotSetting": false,
		}
		jsonEnvs = append(jsonEnvs, env)
	}
	return jsonEnvs
}

func getFile(filepath string) (map[string]interface{}, error) {
	yamlData := make(map[string]interface{})
	yamlFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Can't read the file")
	}
	err = yaml.Unmarshal(yamlFile, &yamlData)
	return yamlData, err
}
