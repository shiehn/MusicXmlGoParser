package xmlparser

import (
	"strings"
)

type PitchConvert struct {
	pitch string
}

func(p *PitchConvert) convert() string{
	if p == nil {
		return "0"
	}

	switch strings.ToLower(p.pitch) {
	case "a":
		return "1"
	case "b":
		return "2"
	case "c":
		return "3"
	case "d":
		return "4"
	case "e":
		return "5"
	case "f":
		return "6"
	case "g":
		return "7"
	}

	panic("pitch convert case not found")
}
