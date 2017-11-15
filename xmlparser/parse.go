package xmlparser

import (
	"fmt"
)

func GetBarDuration(musicXML MXLDoc) int {
	return musicXML.Parts[0].Bars[0].Forward.Duration
}

func Parse(musicXML MXLDoc){
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
}

