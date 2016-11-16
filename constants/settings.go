package constants

import (
	"os"
	"strconv"
	"strings"
)

// PortNum is the port where this server runs
var PortNum = 3001

// LoadEnvironmentSettings loads settings from environment variables
func LoadEnvironmentSettings() {
}

func loadBooleanSetting(setting *bool, envKey string) {
	envVal := os.Getenv(envKey)
	if strings.Compare(envVal, "true") == 0 {
		*setting = true
	} else if strings.Compare(envVal, "false") == 0 {
		*setting = false
	}
}

func loadStringSetting(setting *string, envKey string) {
	envVal := os.Getenv(envKey)
	if len(envVal) > 0 {
		*setting = envVal
	}
}

func loadIntSetting(setting *int, envKey string) {
	envVal := os.Getenv(envKey)
	if len(envVal) > 0 {
		val, err := strconv.Atoi(envVal)
		if err == nil {
			*setting = val
		}
	}
}
