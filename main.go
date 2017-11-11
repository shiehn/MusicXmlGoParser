package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"github.com/MusicXmlGoParser/xmlparser"
)

var personXML = []byte(`
    `)

type Person struct {
	Name string `xmlparser:"name"`
	Addresses []struct {
		Street string `xmlparser:"street"`
		City   string `xmlparser:"city"`
		Type   string `xmlparser:"type,attr"`
	} `xmlparser:"addresses>address"`
}

func main() {

	xmlparser.Read()
	xmlparser.JustATest()

	personXML, err := ioutil.ReadFile("test.xmlparser")
	if err != nil {
		fmt.Print("XML READ ERROR!!!")
		}

	var luann Person
	xml.Unmarshal(personXML, &luann)
	fmt.Println(luann)
}