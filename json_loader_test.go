package widecfg_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamillosantos/widecfg"
	"github.com/jamillosantos/widecfg/testingutils"
)

var _ = Describe("JSONLoader", func() {
	It("should load a json", func() {
		buff := testingutils.NewBufferReader()
		fmt.Fprint(buff, `{
  "prop1": "value1",
  "prop2": 2,
  "prop3": {
    "prop4": true
  }
}`)
		defer buff.Close()

		configMap := &widecfg.ConfigMap{}
		config := widecfg.NewConfig(configMap)
		loader := &widecfg.JSONLoader{}
		Expect(loader.Load(configMap, buff)).NotTo(HaveOccurred())
		Expect(config.GetString("prop1")).To(Equal("value1"))
		Expect(config.GetInt("prop2")).To(Equal(2))
		Expect(config.GetBool("prop3.prop4")).To(BeTrue())
	})

	It("should fail loading a json", func() {
		buff := testingutils.NewBufferReader()
		fmt.Fprint(buff, `{
  "prop1": "value1",
`)
		defer buff.Close()

		configMap := &widecfg.ConfigMap{}
		loader := &widecfg.JSONLoader{}
		Expect(loader.Load(configMap, buff)).To(HaveOccurred())
	})
})
