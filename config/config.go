package config

import (
	"fmt"
	"strconv"
	"strings"
)

const c_EXTENSION_CONFIG_PREFIX = "reports.tagtree."
const c_HIERARCHY_CONFIG_PREFIX = c_EXTENSION_CONFIG_PREFIX + "hierarchy."
const c_MAX_DEPTH_CONFIG_KEY = c_EXTENSION_CONFIG_PREFIX + "maxdepth"
const PRUNE_CONFIG_KEY = c_EXTENSION_CONFIG_PREFIX + "prune"

type ConfigurationKeyNotFound struct {
	Key string
}

func (e *ConfigurationKeyNotFound) Error() string {
	return fmt.Sprintf("Configuration key [%s] not found", e.Key)
}

func ExtractInitialHierarchy(rawConfig map[string]string) map[string][]string {
	hierachy := make(map[string][]string, 0)
	for key, value := range rawConfig {
		if strings.HasPrefix(key, c_HIERARCHY_CONFIG_PREFIX) {
			hierachy[strings.TrimPrefix(key, c_HIERARCHY_CONFIG_PREFIX)] = strings.Fields(value)
		}
	}
	return hierachy
}

func IsDebugEnabled(rawConfig map[string]string) bool {
	debug := rawConfig["debug"]
	return debug == "on" || debug == "yes" || debug == "true" || debug == "y" || debug == "1"
}

func RetrieveMaxDepth(config map[string]string) (uint64, error) {
	if stringMaxDepth, ok := config[c_MAX_DEPTH_CONFIG_KEY]; ok {
		return strconv.ParseUint(stringMaxDepth, 10, 8)
	} else {
		return 0, &ConfigurationKeyNotFound{Key: c_MAX_DEPTH_CONFIG_KEY}
	}
}
