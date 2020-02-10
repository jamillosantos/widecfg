package widecfg

type Getter interface {
	Get(key string) (interface{}, bool)
}
