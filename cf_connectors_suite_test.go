package cf_connectors_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCfConnectors(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CF Connectors Suite")
}
