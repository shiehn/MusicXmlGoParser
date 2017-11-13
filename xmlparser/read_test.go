package xmlparser_test

import (
	. "github.com/MusicXmlGoParser/xmlparser"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Read", func() {

	Context("Just a test", func() {
		It("should be mystring", func() {
			Expect(JustATest()).To(Equal("mystring"))
		})
	})
})
