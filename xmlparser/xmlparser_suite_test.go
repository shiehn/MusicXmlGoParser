package xmlparser_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestXmlparser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Xmlparser Suite")
}
