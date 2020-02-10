package widecfg

import (
	"github.com/joho/godotenv"
)

type EnvFileLoader struct{}

func (envFileLoader *EnvFileLoader) Load(reader Reader) (*ConfigMap, error) {
	envMap, err := godotenv.Parse(reader)
	if err != nil {
		return nil, err
	}
	cm := make(ConfigMap, len(envMap))
	for k, v := range envMap {
		cm[k] = v
	}
	return &cm, nil
}
