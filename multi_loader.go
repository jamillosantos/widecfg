package widecfg

type MultiLoader struct {
	reader  Reader
	loaders []Loader
	getter  func(Getter) Getter
}

func NewMultiLoader() *MultiLoader {
	return &MultiLoader{
		loaders: make([]Loader, 0, 1),
	}
}

func (loader *MultiLoader) File(fName string) *MultiLoader {
	loader.reader = NewFileReader(fName)
	return loader
}

func (loader *MultiLoader) Env(prefix string) *MultiLoader {
	loader.getter = func(getter Getter) Getter {
		return NewEnvGetter(prefix, getter)
	}
	return loader
}

func (loader *MultiLoader) JSON() *MultiLoader {
	loader.loaders = append(loader.loaders, &JSONLoader{})
	return loader
}

func (loader *MultiLoader) HCL() *MultiLoader {
	loader.loaders = append(loader.loaders, &HCLLoader{})
	return loader
}

func (loader *MultiLoader) TOML() *MultiLoader {
	loader.loaders = append(loader.loaders, &TOMLLoader{})
	return loader
}

func (loader *MultiLoader) YAML() *MultiLoader {
	loader.loaders = append(loader.loaders, &YAMLLoader{})
	return loader
}

func (loader *MultiLoader) Load() (*Config, error) {
	defer loader.reader.Close()
	configMap := make(ConfigMap)
	var err error
	for _, ldr := range loader.loaders {
		err = ldr.Load(&configMap, loader.reader)
		if err != nil {
			return nil, err
		}
	}
	var getter Getter = &configMap
	if loader.getter != nil {
		getter = loader.getter(getter)
	}
	return NewConfig(getter), nil
}
