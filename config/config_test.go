package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConfigurationKeyNotFoundError(t *testing.T) {
	err := ConfigurationKeyNotFound{
		Key: "key",
	}
	if err.Error() != "Configuration key [key] not found" {
		t.Errorf("Error message for ConfigurationKeyNotFound not correct: %s", err.Error())
	}
}

func TestExtractInitialHierarchy(t *testing.T) {
	hierarchy := ExtractInitialHierarchy(map[string]string{
		"reports.tagtree.hierarchy.1": "tag1 tag11",
		"some.other.config":           "value",
		"reports.tagtree.hierarchy.2": "tag2 tag21",
	})
	expected := map[string][]string{
		"1": {"tag1", "tag11"},
		"2": {"tag2", "tag21"},
	}
	for key, value := range expected {
		t.Run(key, func(t *testing.T) {
			if !reflect.DeepEqual(hierarchy[key], value) {
				t.Errorf("Hierarchy key [%s] has value [%s] instead of expected value [%s]", key, hierarchy[key], value)
			}
		})
	}
}

func TestIsDebugEnabled(t *testing.T) {
	tests := map[string]bool{
		"on":   true,
		"yes":  true,
		"true": true,
		"y":    true,
		"1":    true,
	}
	for value, expected := range tests {
		t.Run(fmt.Sprintf("%s-%t", value, expected), func(t *testing.T) {
			if IsDebugEnabled(map[string]string{"debug": value}) != expected {
				t.Errorf("Test with value [%s] was not [%t]", value, expected)
			}
		})
	}
}

func TestRetrieveMaxDepth(t *testing.T) {
	config := map[string]string{
		c_MAX_DEPTH_CONFIG_KEY: "12",
	}
	expected := uint64(12)
	if maxDepth, err := RetrieveMaxDepth(config); err == nil {
		if maxDepth != expected {
			t.Errorf("Actual maxDepth %d different from expected %d", maxDepth, expected)
		}
	} else {
		t.Fatal(err)
	}
}

func TestRetrieveMaxDepthNoConfigKey(t *testing.T) {
	config := map[string]string{}
	if _, err := RetrieveMaxDepth(config); err != nil {
		switch err.(type) {
		case *ConfigurationKeyNotFound:
			break
		default:
			t.Error("Method didn't return a ConfigurationKeyNotFound error")
		}
	} else {
		t.Error("Method didn't return error when configuration key is missing")
	}
}

func TestRetrieveMaxDepthParseError(t *testing.T) {
	config := map[string]string{
		c_MAX_DEPTH_CONFIG_KEY: "not a number",
	}
	if _, err := RetrieveMaxDepth(config); err == nil {
		t.Error("Method didn't return error when configuration key is not a parsable unsigned int")
	}
}
