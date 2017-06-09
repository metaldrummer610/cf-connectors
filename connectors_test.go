package cf_connectors_test

import (
	"io/ioutil"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connectors", func() {
	BeforeEach(func() {
		rawConfig, err := ioutil.ReadFile("test_environment_variables.json")
		Expect(err).NotTo(HaveOccurred())
		os.Setenv("VCAP_SERVICES", string(rawConfig))
	})

	It("should parse all services from VCAP Env", func() {

	})
})
