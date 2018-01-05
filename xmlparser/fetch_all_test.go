package xmlparser_test

import (
	"github.com/MusicXmlGoParser/xmlparser"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FetchAll", func() {
	Context("Should", func() {
		It("return list of filenames", func() {
			Expect(len(xmlparser.FetchAll())).ShouldNot(Equal(0))
		})
	})
})
