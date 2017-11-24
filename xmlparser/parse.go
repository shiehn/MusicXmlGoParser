package xmlparser

import (
	"fmt"
	"strings"
	"strconv"
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

func ParseChordsFromBar(musicXML MXLDoc, index int) string {
	var chords string
	bar := musicXML.Parts[0].Bars[index + 1]
	for i, tag := range bar.Harmonies {
			if tag.Print == "" {

				sharpFlat := "_"
					if bar.Harmonies[i].Root.SharpFlat == 1 {
						sharpFlat = "s"
					}else if bar.Harmonies[i].Root.SharpFlat == -1 {
						sharpFlat = "b"
					}

				chords = chords + bar.Harmonies[i].Root.RootNote + sharpFlat + "-" + bar.Harmonies[i].Type + "-"
				}
	}
	return chords
}

func ParseNotesFormBar(musicXML MXLDoc, index int) string {
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
			} else {
				sharpFlat = "_"
			}

			octave := strconv.Itoa(note.Pitch.Octave)
			notes = notes + note.Pitch.Step + sharpFlat + octave +  "-" + note.Type + "-"
		}
	}

	return notes
}

func Parse(musicXML MXLDoc) string {
	chordsAndNotes := ""
	barCount := GetBarCount(musicXML)
	for i := 0; i < barCount; i++ {
		chordsAndNotes = chordsAndNotes + ParseChordsFromBar(musicXML, i)
		chordsAndNotes = chordsAndNotes + ParseNotesFormBar(musicXML, i)
	}

	return chordsAndNotes
}

