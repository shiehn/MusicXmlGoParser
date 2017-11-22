package xmlparser

import (
	"fmt"
	"strings"
)

func GetBarDuration(musicXML MXLDoc) int {
	return musicXML.Parts[0].Bars[0].Forward.Duration
}

func GetSixteenthNote(musicXML MXLDoc) int {
	return musicXML.Parts[0].Bars[0].Forward.Duration/16
}

func GetBarCount(musicXML MXLDoc) int {
	return len(musicXML.Parts[0].Bars)-1
}

func GetChords(musicXML MXLDoc, index int) string {
	var chords string
	bar := musicXML.Parts[0].Bars[index + 1]
	for i, tag := range bar.Harmonies {
			if tag.Print == "" {
				chords = chords + bar.Harmonies[i].Root.Pitch + "-" + bar.Harmonies[i].Type + "-"
				}
	}
	return chords
}

func ParseBar(musicXML MXLDoc, index int) string {
	index = index + 1
	bar := musicXML.Parts[0].Bars[index]

	var notes string
	for _, note := range bar.Notes {
		/*
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
*/
		if strings.Contains(fmt.Sprintf("%v", note.Rest), "rest") {
			//dur := note.Duration / GetSixteenthNote(musicXML)
			notes = notes + "r-" + note.Type + "-"
		} else {
			//dur := note.Duration / GetSixteenthNote(musicXML)
			sharpFlat := ""
			if note.Pitch.Accidental == -1 {
				sharpFlat = "b"
			} else if note.Pitch.Accidental == 1 {
				sharpFlat = "s"
			}
			notes = notes + note.Pitch.Step + sharpFlat + "-" + note.Type + "-"
		}
	}

	return notes
}

func Parse(musicXML MXLDoc) string {

	notes := ""
	barCount := GetBarCount(musicXML)
	for i := 0; i < barCount; i++ {
		notes = ParseBar(musicXML, i)
	}

	return notes

/*
	for barIndex, bar := range musicXML.Parts[0].Bars {
		fmt.Print("%%%%%%%%%%%% MEASURE %%%%%%%%%%%%%% \n")
		fmt.Print("TIME: \n")
		fmt.Print(bar.Atters)
		fmt.Print("\n")

		if len(bar.Harmonies) > 1 {
			fmt.Print("CHORD ONE ---- \n")
			fmt.Print("%v \n", bar.Harmonies[0])
			fmt.Print("CHORD TWO ---- \n")
			fmt.Print("%v \n", bar.Harmonies[1])
		}else {
			panic(fmt.Sprintf("MISSING CHORD!!! ON BAR %v", barIndex))
		}

		for _, note := range bar.Notes {
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
*/
}

