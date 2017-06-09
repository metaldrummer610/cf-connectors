package cf_connectors_test

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "pivotal.io/cf-connectors"
)

var _ = Describe("ENV_VARS", func() {
	var vs VcapServices

	BeforeEach(func() {
		rawConfig, err := ioutil.ReadFile("test_environment_variables.json")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(json.Unmarshal(rawConfig, &vs)).To(Succeed())
	})

	It("should parse correctly", func() {
		Expect(vs).To(HaveLen(3))
		Expect(vs["elephantsql"]).NotTo(BeNil())
		Expect(vs["cloudamqp"]).NotTo(BeNil())
		Expect(vs["rediscloud"]).NotTo(BeNil())
	})
})
