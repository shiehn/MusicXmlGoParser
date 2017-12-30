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
		var parser Parser
		BeforeEach(func() {
			musicXML, err := ioutil.ReadFile(FOUR_BAR_ASSETS)
			if err != nil {
				panic(err)
			}
			xml.Unmarshal(musicXML, &xmlDoc)

			parser = Parser{
				MusicXml: xmlDoc,
			}
		})

		It("should parse notes from a bar", func() {
			//			Expect(parser.ParseNotesFormBar(0)).To(Equal("r-eighth-dot-D_4-16th-nodot-A_4-eighth-dot-Ab4-16th-nodot-Ab4-quarter-dot-G_4-eighth-nodot-"))
			output := parser.ParseNotesFormBar(0)
			expected := "0000-0000-0000-4140-1140-1141-1141-1040-1040-1041-1041-1041-1041-1041-7140-7141"
			fmt.Println(output)
			fmt.Println(expected)
			Expect(output).To(Equal(expected))




			//"61-414-515-0000-0000-0000-4140-1140-1141-1141-1040-1040-1041-1041-1041-1041-1041-7140-7141-61-411-205-6140-6141-6141-6141-4140-4141-6140-6141-6140-6141-6141-6141-6141-6141-6141-6141"
            //"61-411-205-6140-6141-6141-6141-4140-4141-6140-6141-6140-6141-6141-6141-6141-6141-6141-6141"






			//61-414-515-
			//0000-0000-0000-4140-1140-1141-1141-1040-1040-1041-1041-1040-1041-1041-1041-1041-1041-7140
		})

		FIt("should concatinate CHORDS AND NOTES", func() {
			output,_ := parser.Parse()
			expected := "61-216-515-0000-0000-0000-4140-1140-1141-1141-1040-1040-1041-1041-1041-1041-1041-7140-7141-61-411-205-6140-6141-6141-6141-4140-4141-6140-6141-6140-6141-6141-6141-6141-6141-6141-6141"
			fmt.Println("************ACTUAL:")
			fmt.Println(output)
			fmt.Println("************EXPECTED:")
			fmt.Println(expected)
			Expect(output).To(ContainSubstring(expected))

			//Expect(parser.Parse()).To(Equal("D_-minor-seventh-E_-dominant-r-eighth-dot-D_4-16th-nodot-A_4-eighth-dot-Ab4-16th-nodot-Ab4-quarter-dot-G_4-eighth-nodot-D_-minor-Bb-dominant-F_4-quarter-nodot-D_4-eighth-nodot-F_4-eighth-nodot-F_4-half-nodot-D_-minor-seventh-B_-minor-seventh-r-eighth-dot-D_4-16th-nodot-F_4-eighth-dot-A_4-16th-nodot-E_5-eighth-dot-E_5-16th-nodot-D_5-eighth-nodot-A_4-eighth-nodot-E_-minor-seventh-A_-dominant-A_4-whole-nodot-"))
		})

		Context("with duration", func() {
			var xmlDoc MXLDoc
			var parser Parser

			BeforeEach(func() {
				musicXML, err := ioutil.ReadFile(FOUR_BAR_ASSETS)
				if err != nil {
					panic(err)
					fmt.Print("XML READ ERROR!!!")
				}
				xml.Unmarshal(musicXML, &xmlDoc)

				parser = Parser{
					MusicXml: xmlDoc,
				}
			})

			It("should have bar length of 160", func() {
				Expect(parser.GetBarDuration()).To(Equal(160))
			})

			It("should have 16th note duration of 10", func() {
				Expect(parser.GetSixteenthNote()).To(Equal(10))
			})

			It("should have 4 bars not including first", func() {
				Expect(parser.GetBarCount()).To(Equal(4))
			})

			It("should return correct bar 1 chords", func() {
				Expect(parser.ParseChordsFromBar(0)).To(Equal("D_-minor-seventh-E_-dominant-"))
			})

			It("should return correct bar 2 chords", func() {
				Expect(parser.ParseChordsFromBar(1)).To(Equal("D_-minor-Bb-dominant-"))
			})

			It("should return correct bar 3 chords", func() {
				Expect(parser.ParseChordsFromBar(2)).To(Equal("D_-minor-seventh-B_-minor-seventh-"))
			})

			It("should return correct bar 4 chords", func() {
				Expect(parser.ParseChordsFromBar(3)).To(Equal("E_-minor-seventh-A_-dominant-"))
			})
		})
	})

	Context("When chords are missing", func() {
		var xmlDoc MXLDoc
		var parser Parser
		BeforeEach(func() {
			musicXML, err := ioutil.ReadFile("C:\\gocode\\src\\github.com\\MusicXmlGoParser\\testassets\\missing_chords.xml")
			if err != nil {
				panic(err)
				fmt.Print("XML READ ERROR!!!")
			}
			xml.Unmarshal(musicXML, &xmlDoc)
			parser = Parser{
				MusicXml: xmlDoc,
			}
		})

		It("should panic", func() {
			_, err := parser.ParseChordsFromBar(0)
			Eventually(err).Should(HaveOccurred())
		})
	})

})
