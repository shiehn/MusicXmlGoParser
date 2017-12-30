package xmlparser

type ChordConvert struct {
	chordType string
}

func (cc *ChordConvert) convert() string {

	switch cc.chordType {
	case "major":
		return "0"
	case "minor":
		return "1"
	case "diminished":
		return "2"
	case "major-seventh":
		return "3"
	case "minor-seventh":
		return "4"
	case "dominant":
		return "5"
		return "6"
	case "minor-seven-flat-five??????":
		return "7"
	}

	panic("could not match chord type in ChordConvert")
}
