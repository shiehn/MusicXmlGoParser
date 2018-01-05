package xmlparser

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	MusicXml MXLDoc
}

func (p *Parser) GetBarDuration() int {
	return p.MusicXml.Parts[0].Bars[0].Forward.Duration
}

func (p *Parser) GetSixteenthNote() int {
	return p.MusicXml.Parts[0].Bars[0].Forward.Duration / 16
}

func (p *Parser) GetBarCount() int {
	return len(p.MusicXml.Parts[0].Bars) - 1
}

func (p *Parser) ParseChordsFromBar(index int) (string, error) {
	var chords []string
	bar := p.MusicXml.Parts[0].Bars[index+1]
	for i, tag := range bar.Harmonies {
		if tag.Print == "" {
			sharpFlat := "1"
			if bar.Harmonies[i].Root.SharpFlat == 1 {
				sharpFlat = "2"
			} else if bar.Harmonies[i].Root.SharpFlat == -1 {
				sharpFlat = "0"
			}

			pc := PitchConvert{pitch: bar.Harmonies[i].Root.RootNote}
			cc := ChordConvert{chordType: bar.Harmonies[i].Type, degree: &bar.Harmonies[i].Degree}
			chordStr := pc.convert() + sharpFlat + cc.convert()
			chords = append(chords, chordStr)
		}
	}

	if len(chords) != 2 {
		errorMessage := fmt.Sprintf("bar %v should have exactly 2 chords but has %v", index, len(chords))
		return "", errors.New(errorMessage)
	}

	return strings.Join(chords, "*"), nil
}

func (p *Parser) ParseNotesFormBar(index int) string {

	index = index + 1
	bar := p.MusicXml.Parts[0].Bars[index]

	var notes string
	for _, note := range bar.Notes {
		dotted := strings.Contains(fmt.Sprintf("%v", note.Dot), "dot")
		duration := getDuration(note.Type, dotted)

		if strings.Contains(fmt.Sprintf("%v", note.Rest), "rest") {
			for i := 0; i < duration; i++ {
				notes = notes + "0000-"
			}
		} else {
			sharpFlat := "1"
			if note.Pitch.Accidental == -1 {
				sharpFlat = "0"
			} else if note.Pitch.Accidental == 1 {
				sharpFlat = "2"
			}

			for i := 0; i < duration; i++ {
				lifeCycle := "1"
				if i == 0 {
					lifeCycle = "0"
				}

				octave := strconv.Itoa(note.Pitch.Octave)
				pitchConv := PitchConvert{pitch: note.Pitch.Step}
				notes = notes + pitchConv.convert() + sharpFlat + octave + lifeCycle + "-"
			}
		}
	}
	return strings.TrimSuffix(notes, "-")
}

func (p *Parser) Parse() (string, error) {
	barCount := p.GetBarCount()
	barCount = p.TotalDivisibleLength(barCount)

	validator := Validate{
		Bars: p.MusicXml.Parts[0].Bars,
	}

	key := KeyConvert{
		Key: p.MusicXml.Parts[0].Bars[0].Atters.Key.Fifths,
	}

	err := validator.CheckDurations()
	if err != nil {
		return "", err
	}

	chordsAndNotes := ""
	for j := 0; j < (barCount/4); j++ {
		chordsAndNotes = chordsAndNotes + fmt.Sprintf(key.convert()) + "^"

		chords := ""
		for i := 0; i < 4; i++ {
			newChords, err := p.ParseChordsFromBar(i+(4*j))
			if err != nil {
				return "", err
			}

			chords = chords + newChords + "*"
		}

		notes := ""
		for i := 0; i < 4; i++ {
			newNotes := p.ParseNotesFormBar(i + (4 * j))
			notes = notes + newNotes
			notes = strings.TrimSuffix(notes, "-")
			if i != 3 {
				notes = notes + "+"
			}
		}

		chordsAndNotes = chordsAndNotes + chords + notes + "#"
	}

	return chordsAndNotes, nil
}

func (p *Parser) TotalDivisibleLength(length int) int {
	result := length - (length%4)
	if result < 0 {
		return 0
	}

	return result
}