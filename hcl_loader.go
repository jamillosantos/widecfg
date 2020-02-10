package widecfg

import (
	"io/ioutil"

	"github.com/hashicorp/hcl"
)

type HCLLoader struct{}

func (yamlReader *HCLLoader) Load(config *ConfigMap, reader Reader) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	cm := make(map[string]interface{}, 0)
	err = hcl.Unmarshal(data, &cm)
	if err != nil {
		return err
	}
	for k, v := range cm {
		(*config)[k] = v
	}
	return nil
}
