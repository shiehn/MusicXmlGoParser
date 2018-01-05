package xmlparser

import (
	"encoding/xml"
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var (
	FOUR_BAR_ASSETS string = "C:\\gocode\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_four_bars.xml"
	ONE_BAR_ASSETS  string = "C:\\gocode\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_one_bar.xml"
)

var _ bool = Describe("Validate", func() {
	Context("When durations are correct", func() {
		var xmlDoc MXLDoc
		BeforeEach(func() {
			musicXML, err := ioutil.ReadFile(FOUR_BAR_ASSETS)
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}
			xml.Unmarshal(musicXML, &xmlDoc)
		})

		It("should panic if empty bars passesing", func() {

		})

		It("should not panic", func() {
			v := Validate{
				Bars: xmlDoc.Parts[0].Bars,
			}

			Expect(v.CheckDurations()).ToNot(HaveOccurred())
		})
	})
})
