package widecfg_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamillosantos/widecfg"
	"github.com/jamillosantos/widecfg/testingutils"
)

var _ = Describe("TOMLLoader", func() {
	It("should load a toml", func() {
		buff := testingutils.NewBufferReader()
		fmt.Fprint(buff, `
prop1 = "value1"
prop2 = 2
[prop3]
  prop4 = true
`)
		defer buff.Close()

		configMap := &widecfg.ConfigMap{}
		config := widecfg.NewConfig(configMap)
		loader := &widecfg.TOMLLoader{}
		Expect(loader.Load(configMap, buff)).NotTo(HaveOccurred())
		Expect(config.GetString("prop1")).To(Equal("value1"))
		Expect(config.GetInt("prop2")).To(Equal(2))
		Expect(config.GetBool("prop3.prop4")).To(BeTrue())
	})

	It("should fail loading a toml", func() {
		buff := testingutils.NewBufferReader()
		fmt.Fprint(buff, `this is an invalid TOML file`)
		defer buff.Close()

		configMap := &widecfg.ConfigMap{}
		loader := &widecfg.TOMLLoader{}
		Expect(loader.Load(configMap, buff)).To(HaveOccurred())
	})
})
