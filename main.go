package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"github.com/MusicXmlGoParser/xmlparser"
)

var musicXML = []byte(`
    `)


// MXLDoc holds all data for a music xml file
type MXLDoc struct {
	Score          xml.Name `xml:"score-partwise"`
	Identification `xml:"identification"`
	Parts          []Part `xml:"part"`
}

// Identification holds all of the ident information for a music xml file
type Identification struct {
	Composer string `xml:"creator"`
	Encoding `xml:"encoding"`
	Rights   string `xml:"rights"`
	Source   string `xml:"source"`
	Title    string `xml:"movement-title"`
}

// Encoding holds encoding info
type Encoding struct {
	Software string `xml:"software"`
	Date     string `xml:"encoding-date"`
}

// Part represents a part in a piece of music
type Part struct {
	Id       string    `xml:"id,attr"`
	Measures []Measure `xml:"measure"`
}

// Measure represents a measure in a piece of music
type Measure struct {
	Number int        `xml:"number,attr"`
	Atters Attributes `xml:"attributes"`
	Harmonies []Harmony `xml:"harmony"`
	Notes  []Note     `xml:"note"`
}

// Attributes represents
type Attributes struct {
	Key       Key  `xml:"key"`
	Time      Time `xml:"time"`
	Divisions int  `xml:"divisions"`
	Clef      Clef `xml:"clef"`
}

type Root struct {
	Pitch string `xml:"root-step"`
}

type Harmony struct {
	Root Root `xml:"root"`
	Type string `xml:"kind"`
}

// Clef represents a clef change
type Clef struct {
	Sign string `xml:"sign"`
	Line int    `xml:"line"`
}

// Key represents a key signature change
type Key struct {
	Fifths int    `xml:"fifths"`
	Mode   string `xml:"mode"`
}

// Time represents a time signature change
type Time struct {
	Beats    int `xml:"beats"`
	BeatType int `xml:"beat-type"`
}

// Note represents a note in a measure
type Note struct {
	Pitch    Pitch    `xml:"pitch"`
	Duration int      `xml:"duration"`
	Voice    int      `xml:"voice"`
	Type     string   `xml:"type"`
	Rest     xml.Name `xml:"rest"`
	Chord    xml.Name `xml:"chord"`
	Tie      Tie      `xml:"tie"`
}

// Pitch represents the pitch of a note
type Pitch struct {
	Accidental int8   `xml:"alter"`
	Step       string `xml:"step"`
	Octave     int    `xml:"octave"`
}

// Tie represents whether or not a note is tied.
type Tie struct {
	Type string `xml:"type,attr"`
}

//type Person struct {
//	Name string `xmlparser:"name"`
//	Addresses []struct {
//		Street string `xmlparser:"street"`
//		City   string `xmlparser:"city"`
//		Type   string `xmlparser:"type,attr"`
//	} `xmlparser:"addresses>address"`
//}

func main() {

	xmlparser.Read()
	xmlparser.JustATest()

	musicXML, err := ioutil.ReadFile("C:\\Users\\steve\\Desktop\\SmartScoreScans\\26_angel_eyes.xml")
	if err != nil {
		fmt.Print("XML READ ERROR!!!")
		}

	var xmlDoc MXLDoc
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