package util

import (
	"os"

	"gopkg.in/yaml.v3"
)

func WriteYamlToFile(dst string, data any) error {
	str, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(dst, str, 0644)
	return err
}
