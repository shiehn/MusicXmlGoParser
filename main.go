package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"github.com/MusicXmlGoParser/xmlparser"
	"log"
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

		err = xml.Unmarshal(musicXML, &xmlDoc)
		if err != nil {
			panic(err)
		}

		parser := xmlparser.Parser{
			MusicXml: xmlDoc,
		}

		audioStr, err := parser.Parse()
		if err != nil{
			log.Fatal(err)
		}

		fmt.Printf("OUTPUT: %s", audioStr)
}