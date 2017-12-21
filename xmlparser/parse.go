package xmlparser

import (
	"fmt"
	"strings"
	"strconv"
	"errors"
)

type Parser struct {
	MusicXml MXLDoc
}

func (p *Parser)GetBarDuration() int {
	return p.MusicXml.Parts[0].Bars[0].Forward.Duration
}

func (p *Parser)GetSixteenthNote() int {
	return p.MusicXml.Parts[0].Bars[0].Forward.Duration / 16
}

func (p *Parser)GetBarCount() int {
	return len(p.MusicXml.Parts[0].Bars) - 1
}

func (p *Parser)ParseChordsFromBar(index int) (string, error) {
	var chords []string
	bar := p.MusicXml.Parts[0].Bars[index+1]
	for i, tag := range bar.Harmonies {
		if tag.Print == "" {
			sharpFlat := "_"
			if bar.Harmonies[i].Root.SharpFlat == 1 {
				sharpFlat = "s"
			} else if bar.Harmonies[i].Root.SharpFlat == -1 {
				sharpFlat = "b"
			}

			chordStr := bar.Harmonies[i].Root.RootNote + sharpFlat + "-" + bar.Harmonies[i].Type
			chords = append(chords, chordStr)
		}
	}

	if len(chords) != 2 {
		errorMessage := fmt.Sprintf("bar %v should have exactly 2 chords but has %v", index, len(chords))
		return "", errors.New(errorMessage)

	}

	return strings.Join(chords, "-") + "-", nil
}

func (p *Parser)ParseNotesFormBar(index int) string {
	index = index + 1
	bar := p.MusicXml.Parts[0].Bars[index]

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
		dotted := strings.Contains(fmt.Sprintf("%v", note.Dot), "dot")

		if strings.Contains(fmt.Sprintf("%v", note.Rest), "rest") {
			//dur := note.Duration / GetSixteenthNote(musicXML)
			notes = notes + "r-" + CreateDuration(note.Type, dotted) + "-"
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
			notes = notes + note.Pitch.Step + sharpFlat + octave + "-" + CreateDuration(note.Type, dotted) + "-"
		}
	}

	return notes
}

func CreateDuration(duration string, isDotted bool) string {

	dotted := "nodot"
	if isDotted {
		dotted = "dot"
	}

	return fmt.Sprintf("%s-%s", duration, dotted)
}

func (p *Parser)Parse() (string, error) {

	chordsAndNotes := ""
	barCount := p.GetBarCount()

	validator := Validate{
		Bars: p.MusicXml.Parts[0].Bars,
	}

	err := validator.CheckDurations()
	if err != nil {
		return "", err
	}

	for i := 0; i < barCount; i++ {
	    newChords, err := p.ParseChordsFromBar(i)
	    if err != nil {
	    	return "", err
		}

		chordsAndNotes = chordsAndNotes + newChords
		chordsAndNotes = chordsAndNotes + p.ParseNotesFormBar(i)
	}

	return chordsAndNotes, nil
}

