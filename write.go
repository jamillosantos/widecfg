package widecfg

type Writer interface {
	Write(cfg *Config) error
}
