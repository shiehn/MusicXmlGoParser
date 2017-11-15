package xmlparser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

var _ = Describe("Read", func() {
	Context("Just a test", func() {
		var xmlDoc MXLDoc

		BeforeEach(func(){
			musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_one_bar.xml")
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}

			xml.Unmarshal(musicXML, &xmlDoc)

			fmt.Print(xmlDoc)
		})

		It("should be mystring", func() {
			Expect("mystring").To(Equal("mystring"))
		})

		It("should have bar length of 160", func(){
			Expect(GetBarDuration(xmlDoc)).To(Equal(160))
		})
	})
})
