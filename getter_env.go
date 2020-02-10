package widecfg

import (
	"os"
	"strings"
)

type EnvGetter struct {
	Getter
	prefix string
}

func NewEnvGetter(prefix string, getter Getter) Getter {
	return &EnvGetter{
		Getter: getter,
		prefix: prefix,
	}
}

func (getter *EnvGetter) Get(key string) (interface{}, bool) {
	props := strings.Split(key, ".")
	envName := getter.prefix
	for i, propName := range props {
		if i > 0 {
			envName += "_"
		}
		envName += strings.ToUpper(propName)
	}
	value, ok := os.LookupEnv(envName)
	if !ok {
		return getter.Getter.Get(key)
	}
	return value, true
}
