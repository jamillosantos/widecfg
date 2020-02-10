package widecfg

type Loader interface {
	Load(*ConfigMap, Reader) error
}
