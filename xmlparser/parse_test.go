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
			musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_one_bar.xml")
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}

			xml.Unmarshal(musicXML, &xmlDoc)

			//fmt.Print(xmlDoc)
		})

		FIt("should parse all bars", func() {

			fmt.Println(Parse(xmlDoc))
					Expect(Parse(xmlDoc)).To(Equal("r-eigth-d4-16th-a4-eight-ab4-16th-ab4-quarter-g4-eight"))
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

		It("should return correct bar 1 chords", func(){
			Expect(GetChords(xmlDoc, 0)).To(Equal("d-0-min7 d-1-min7"))
		})

		It("should return correct bar 2 chords", func(){
			Expect(GetChords(xmlDoc, 1)).To(Equal("d-0-min7 d-1-min7"))
		})

		It("should return correct bar 3 chords", func(){
			Expect(GetChords(xmlDoc, 2)).To(Equal("d-0-min7 d-1-min7"))
		})

		It("should return correct bar 4 chords", func(){
			Expect(GetChords(xmlDoc, 3)).To(Equal("d-0-min7 d-1-min7"))
		})
	})
})
