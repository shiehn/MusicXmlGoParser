package xmlparser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"fmt"
	"encoding/xml"
)

var _ = Describe("MxlDoc", func() {
	Context("when parsed", func() {
		var xmlDoc MXLDoc
		BeforeEach(func(){
			musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_one_bar.xml")
			if err != nil {
				panic(err)
			}
			xml.Unmarshal(musicXML, &xmlDoc)
		})

		It("should parse notes from a bar", func() {
				Expect(ParseNotesFormBar(xmlDoc, 0)).To(Equal("r-eighth-D_4-16th-A_4-eighth-Ab4-16th-Ab4-quarter-G_4-eighth-"))
		})

		It("should concatinate CHORDS AND NOTES", func(){
			Expect(Parse(xmlDoc)).To(Equal("D_-minor-seventh-E_-dominant-r-eighth-D_4-16th-A_4-eighth-Ab4-16th-Ab4-quarter-G_4-eighth-"))
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
			Expect(ParseChordsFromBar(xmlDoc, 0)).To(Equal("D_-minor-seventh-E_-dominant-"))
		})

		It("should return correct bar 2 chords", func(){
			Expect(ParseChordsFromBar(xmlDoc, 1)).To(Equal("D_-minor-Bb-dominant-"))
		})

		It("should return correct bar 3 chords", func(){
			Expect(ParseChordsFromBar(xmlDoc, 2)).To(Equal("D_-minor-seventh-B_-minor-seventh-"))
		})

		It("should return correct bar 4 chords", func(){
			Expect(ParseChordsFromBar(xmlDoc, 3)).To(Equal("E_-minor-seventh-A_-dominant-"))
		})
	})
})
