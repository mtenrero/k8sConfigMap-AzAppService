package cmapparser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const testConfigMap = "./_testFiles/configMap.yaml"

func TestReadYamlFile(t *testing.T) {

	parsedConfigMap, err := ParseConfigMapEnvVars(testConfigMap)

	assert.Nil(t, err, "Error parsing YAML ConfigMap")
	assert.NotNil(t, parsedConfigMap, "parsedConfigMap returns an empty object!")
	assert.Len(t, parsedConfigMap, 4, "Error determining YAML objects from parsedConfigMap")
}

func TestYAMLisConfigMap(t *testing.T) {
	var yamlMock = map[string]interface{}{
		"kind": "ConfigMap",
	}
	err := ensureConfigMap(yamlMock)
	assert.Nil(t, err, "It is expected to be a ConfigMap")

	yamlMock["kind"] = "NoConfigMap"
	err = ensureConfigMap(yamlMock)
	assert.NotNil(t, err, "Is expected to not detect a k8s ConfigMap")
}

func TestJSONTransform(t *testing.T) {
	parsedConfigMap, _ := ParseConfigMapEnvVars(testConfigMap)
	jsonObject := ConvertToAzureJSON(parsedConfigMap)

	assert.Len(t, jsonObject, 4)
	assert.Contains(t, jsonObject, map[string]interface{}{"name": "PZ_ENV_COLOR", "slotSetting": false, "value": "#3F51B5"})
}
