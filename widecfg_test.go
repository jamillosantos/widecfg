package widecfg_test

import (
	"testing"

	"github.com/novln/macchiato"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

// TestWideCfg triggers the macchiato for starting the tests.
func TestWideCfg(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	macchiato.RunSpecs(t, "widecfg Test Suite")
}
