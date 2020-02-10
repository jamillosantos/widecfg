package widecfg

import (
	"encoding/json"
)

type JSONLoader struct{}

func (jsonLoader *JSONLoader) Load(config *ConfigMap, reader Reader) error {
	cm := make(map[string]interface{}, 0)
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&cm)
	if err != nil {
		return err
	}
	for k, v := range cm {
		(*config)[k] = v
	}
	return nil
}
