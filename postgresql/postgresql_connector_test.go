package postgresql_test

import (
	"encoding/json"
	"io/ioutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	connectors "pivotal.io/cf-connectors"
	. "pivotal.io/cf-connectors/postgresql"
)

var _ = Describe("Postgresql Connector", func() {
	var config connectors.VcapServices

	BeforeEach(func() {
		rawConfig, err := ioutil.ReadFile("../test_environment_variables.json")
		Expect(err).ShouldNot(HaveOccurred())
		Expect(json.Unmarshal(rawConfig, &config)).To(Succeed())
	})

	It("should parse the configuration properly", func() {
		configs, err := Parse(&config)
		Expect(err).NotTo(HaveOccurred())
		Expect(configs).NotTo(BeNil())
		Expect(configs).To(HaveLen(1))

		config := configs[0]
		Expect(config.Name).To(Equal("pgsql_db"))
		Expect(config.Credentials.Uri).To(Equal("postgres://user:pass@host.com:5432/db"))
	})

})
