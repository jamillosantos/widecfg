package widecfg_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamillosantos/widecfg"
)

var _ = Describe("FileReader", func() {
	It("should load from a file", func() {
		reader := widecfg.NewFileReader("./testdata/config1.json")
		Expect(reader.Open()).To(Succeed())
		defer func() {
			defer GinkgoRecover()
			Expect(reader.Close()).To(Succeed())
		}()

		configMap := &widecfg.ConfigMap{}
		config := widecfg.NewConfig(configMap)
		loader := &widecfg.JSONLoader{}
		Expect(loader.Load(configMap, reader)).NotTo(HaveOccurred())
		Expect(config.GetString("prop1")).To(Equal("value1"))
		Expect(config.GetInt("prop2")).To(Equal(2))
		Expect(config.GetBool("prop3.prop4")).To(BeTrue())
	})

	It("should fail loading a non existing file", func() {
		reader := widecfg.NewFileReader("./testdata/nonexisting_config")
		err := reader.Open()
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("no such file or directory"))
	})

	It("should fail reading a non open file", func() {
		reader := widecfg.NewFileReader("./testdata/nonexisting_config")
		n, err := reader.Read(nil)
		Expect(n).To(BeZero())
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("not open"))
	})
})
