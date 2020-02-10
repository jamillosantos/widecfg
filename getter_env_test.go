package widecfg_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamillosantos/widecfg"
)

var _ = Describe("EnvGetter", func() {
	It("should load from the environment", func() {
		os.Setenv("PROP1", "value1")
		os.Setenv("PROP3_PROP4", "true")
		configMap := &widecfg.ConfigMap{
			"prop2": 2,
		}
		config := widecfg.NewConfig(widecfg.NewEnvGetter("", configMap))
		Expect(config.GetString("prop1")).To(Equal("value1"))
		Expect(config.GetInt("prop2")).To(Equal(2))
		Expect(config.GetBool("prop3.prop4")).To(BeTrue())
	})
})
