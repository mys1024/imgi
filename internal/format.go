package internal

import (
	"bytes"
	"encoding/json"

	"github.com/pelletier/go-toml/v2"
	"gopkg.in/yaml.v3"
)

type Formatter func(scanResult *ScanResult) (string, error)

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

func Toml(scanResult *ScanResult) (string, error) {
	toml, err := toml.Marshal(scanResult)
	if err != nil {
		return "", err
	}
	return string(toml), nil
}

func Json(scanResult *ScanResult) (string, error) {
	buffer := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "  ")
	err := encoder.Encode(scanResult)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
