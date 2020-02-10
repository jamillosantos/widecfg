package widecfg

import (
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type TOMLLoader struct{}

func (tomlReader *TOMLLoader) Load(config *ConfigMap, reader Reader) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	tree, err := toml.LoadBytes(data)
	if err != nil {
		return err
	}
	cm := tree.ToMap()
	for k, v := range cm {
		(*config)[k] = v
	}
	return nil
}
