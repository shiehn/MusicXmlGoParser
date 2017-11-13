package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"github.com/MusicXmlGoParser/xmlparser"
)

var musicXML = []byte(`
    `)

func main() {

	xmlparser.Read()
	xmlparser.JustATest()

	musicXML, err := ioutil.ReadFile("C:\\Users\\steve\\Desktop\\SmartScoreScans\\26_angel_eyes.xml")
	if err != nil {
		fmt.Print("XML READ ERROR!!!")
		}

	var xmlDoc xmlparser.MXLDoc
	xml.Unmarshal(musicXML, &xmlDoc)

	for _, measure := range xmlDoc.Parts[0].Measures {
		fmt.Print("%%%%%%%%%%%% MEASURE %%%%%%%%%%%%%% \n")
		fmt.Print("TIME: \n")
		fmt.Print(measure.Atters)
		fmt.Print("\n")
		fmt.Print("CHORD ---- \n")
		if len(measure.Harmonies) > 0 {
			fmt.Print(measure.Harmonies[0])
		}
		for _, note := range measure.Notes {
			fmt.Print("*************************** \n")
			fmt.Print("\n")
			fmt.Print("PITCH ---- \n")
			fmt.Print(note.Pitch)
			fmt.Print("\n")
			fmt.Print("DURATIon \n")
			fmt.Print(note.Duration)
			fmt.Print("\n")
			fmt.Print("REST \n")
			fmt.Print(note.Rest)
			fmt.Print("\n")
		}
	}

	//fmt.Println(xmlDoc)
}