package widecfg_test

import (
	"time"

	"github.com/jamillosantos/widecfg"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type mismatchedValue struct {
	Value int
}

var _ = Describe("Config", func() {
	Describe("Get", func() {

		It("should get a value that exists", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
			}
			prop1, ok := cfg.Get("prop1")
			Expect(ok).To(BeTrue())
			Expect(prop1).To(Equal("value1"))
		})

		It("should fail getting a value that does not exists", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
			}
			prop1, ok := cfg.Get("prop2")
			Expect(ok).To(BeFalse())
			Expect(prop1).To(BeNil())
		})

		It("should get a value on a subobject", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
				"subprop": map[string]interface{}{
					"subprop1": "subvalue1",
					"subprop2": "subvalue2",
				},
			}
			prop1, ok := cfg.Get("subprop.subprop2")
			Expect(ok).To(BeTrue())
			Expect(prop1).To(Equal("subvalue2"))
		})

		It("should fail getting value from a subobject", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
				"subprop": map[string]interface{}{
					"subprop1": "subvalue1",
					"subprop2": "subvalue2",
				},
			}
			prop1, ok := cfg.Get("subprop.subprop3")
			Expect(ok).To(BeFalse())
			Expect(prop1).To(BeNil())
		})
	})

	Describe("GetString", func() {
		It("should get a string from a existing string value", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
			}
			prop1, err := cfg.GetString("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal("value1"))
		})

		It("should get a string from a existing string value pointer", func() {
			s := "value1"
			cfg := widecfg.Config{
				"prop1": &s,
			}
			prop1, err := cfg.GetString("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal("value1"))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": "value1",
			}
			prop1, err := cfg.GetString("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(""))
		})

		It("should fail to get a string from a existing non string value", func() {
			cfg := widecfg.Config{
				"prop1": 1,
			}
			prop1, err := cfg.GetString("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(""))
		})
	})

	Describe("GetInt", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": 1,
			}
			prop1, err := cfg.GetInt("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(1))
		})

		It("should get the value from a pointer", func() {
			i := 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetInt("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(123))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetInt("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(1))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetInt("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(123))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetInt("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(0))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetInt("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(0))
		})
	})

	Describe("GetUint", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": uint(1),
			}
			prop1, err := cfg.GetUint("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint(1)))
		})

		It("should get the value from a pointer", func() {
			var i uint = 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetUint("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint(123)))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetUint("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint(1)))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetUint("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint(123)))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetUint("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(uint(0)))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetUint("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(uint(0)))
		})
	})

	Describe("GetInt64", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": int64(1),
			}
			prop1, err := cfg.GetInt64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(int64(1)))
		})

		It("should get the value from a pointer", func() {
			var i int64 = 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetInt64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(int64(123)))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetInt64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(int64(1)))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetInt64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(int64(123)))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetInt64("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(int64(0)))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetInt64("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(int64(0)))
		})
	})

	Describe("GetUint64", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": uint64(1),
			}
			prop1, err := cfg.GetUint64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint64(1)))
		})

		It("should get the value from a pointer", func() {
			var i uint64 = 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetUint64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint64(123)))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetUint64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint64(1)))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetUint64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(uint64(123)))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetUint64("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(uint64(0)))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetUint64("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(uint64(0)))
		})
	})

	Describe("GetFloat32", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": float32(1),
			}
			prop1, err := cfg.GetFloat32("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float32(1)))
		})

		It("should get the value from a pointer", func() {
			var i float32 = 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetFloat32("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float32(123)))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetFloat32("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(float32(0)))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetFloat32("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float32(1)))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetFloat32("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float32(123)))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetFloat32("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(float32(0)))
		})
	})

	Describe("GetFloat64", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": float64(1),
			}
			prop1, err := cfg.GetFloat64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float64(1)))
		})

		It("should get the value from a pointer", func() {
			var i float64 = 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetFloat64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float64(123)))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": 123,
			}
			prop1, err := cfg.GetFloat64("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(Equal(float64(0)))
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetFloat64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float64(1)))
		})

		It("should get the value from a string pointer", func() {
			i := "123"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetFloat64("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(float64(123)))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetFloat64("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(Equal(float64(0)))
		})
	})

	Describe("GetBool", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": true,
			}
			prop1, err := cfg.GetBool("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTrue())
		})

		It("should get the value from a pointer", func() {
			var i bool = true
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetBool("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTrue())
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "1",
			}
			prop1, err := cfg.GetBool("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTrue())
		})

		It("should get the value from a string pointer", func() {
			i := "1"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetBool("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTrue())
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{
				"prop1": true,
			}
			prop1, err := cfg.GetBool("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(BeFalse())
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetBool("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(BeFalse())
		})
	})

	Describe("GetTime", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": time.Now(),
			}
			prop1, err := cfg.GetTime("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should get the value from a pointer", func() {
			i := time.Now()
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetTime("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{}
			prop1, err := cfg.GetTime("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1.IsZero()).To(BeTrue())
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": time.Now().Format(widecfg.DefaultTimeFormat),
			}
			prop1, err := cfg.GetTime("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should get the value from a string pointer", func() {
			i := time.Now().Format(widecfg.DefaultTimeFormat)
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetTime("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(BeTemporally("~", time.Now(), time.Second))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetTime("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1.IsZero()).To(BeTrue())
		})
	})

	Describe("GetDuration", func() {
		It("should get the value", func() {
			cfg := widecfg.Config{
				"prop1": time.Second * 123,
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a pointer", func() {
			i := time.Second * 123
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should fail to get a non existing key", func() {
			cfg := widecfg.Config{}
			prop1, err := cfg.GetDuration("prop2")
			Expect(err).To(Equal(widecfg.ErrKeyNotFound))
			Expect(prop1).To(BeZero())
		})

		It("should get the value from a string", func() {
			cfg := widecfg.Config{
				"prop1": "123s",
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a string pointer", func() {
			i := "123s"
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a integer", func() {
			cfg := widecfg.Config{
				"prop1": int(123e9),
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a integer pointer", func() {
			i := int(123e9)
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a int64", func() {
			cfg := widecfg.Config{
				"prop1": int64(123e9),
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should get the value from a int64 pointer", func() {
			i := int64(123e9)
			cfg := widecfg.Config{
				"prop1": &i,
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).ToNot(HaveOccurred())
			Expect(prop1).To(Equal(time.Second * 123))
		})

		It("should fail getting the value due to incompatible data type stored", func() {
			cfg := widecfg.Config{
				"prop1": mismatchedValue{},
			}
			prop1, err := cfg.GetDuration("prop1")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(widecfg.ErrValueWrongType))
			Expect(prop1).To(BeZero())
		})
	})
})
