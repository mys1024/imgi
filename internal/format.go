package internal

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

func Yaml(scanResult *ScanResult) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	encoder := yaml.NewEncoder(buffer)
	encoder.SetIndent(2)
	err := encoder.Encode(scanResult)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buffer.String(), nil
}
