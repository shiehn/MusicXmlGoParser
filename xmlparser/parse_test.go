package xmlparser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

var _ = Describe("MxlDoc", func() {
	Context("Just a test", func() {
		var xmlDoc MXLDoc
		BeforeEach(func(){
			musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_one_angel_eyes.xml")
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}

			xml.Unmarshal(musicXML, &xmlDoc)

			//fmt.Print(xmlDoc)
		})

		It("should parse all bars", func() {
			Expect(Parse(xmlDoc)).To(Equal("dddddddddddd"))
		})
	})

	Context("with duration", func(){
		var xmlDoc MXLDoc
		BeforeEach(func(){
			musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_four_bars.xml")
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}

			xml.Unmarshal(musicXML, &xmlDoc)

			//fmt.Print(xmlDoc)
		})

		It("should have bar length of 160", func(){
			Expect(GetBarDuration(xmlDoc)).To(Equal(160))
		})

		It("should have 16th note duration of 10", func(){
			Expect(GetSixteenthNote(xmlDoc)).To(Equal(10))
		})

		It("should have 4 bars not including first", func(){
			Expect(GetBarCount(xmlDoc)).To(Equal(4))
		})
	})
})
