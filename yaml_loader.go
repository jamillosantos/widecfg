package widecfg

import (
	"gopkg.in/yaml.v3"
)

type YAMLLoader struct{}

func (yamlReader *YAMLLoader) Load(config *ConfigMap, reader Reader) error {
	cm := make(ConfigMap, 0)
	decoder := yaml.NewDecoder(reader)
	err := decoder.Decode(cm)
	if err != nil {
		return err
	}
	for k, v := range cm {
		(*config)[k] = v
	}
	return nil
}
