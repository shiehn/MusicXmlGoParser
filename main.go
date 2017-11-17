package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"github.com/MusicXmlGoParser/xmlparser"
)

var musicXML = []byte(`
    `)

func main() {
	var xmlDoc xmlparser.MXLDoc
		musicXML, err := ioutil.ReadFile("C:\\GoWorkspace\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_four_bars.xml")
		if err != nil {
			panic(err)
			fmt.Print("XML READ ERROR!!!")
		}

		xml.Unmarshal(musicXML, &xmlDoc)

		xmlparser.GetChords(xmlDoc, 2)

		//fmt.Print(xmlDoc)
}