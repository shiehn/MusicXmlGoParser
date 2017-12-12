package main

import (
	"io/ioutil"
	"fmt"
	"encoding/xml"
	"log"
	"github.com/MusicXmlGoParser/xmlparser"
)

var musicXML = []byte(`
    `)

func main() {
	var xmlDoc xmlparser.MXLDoc
		musicXML, err := ioutil.ReadFile("C:\\gocode\\src\\github.com\\MusicXmlGoParser\\testassets\\asset_four_bars.xml")
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