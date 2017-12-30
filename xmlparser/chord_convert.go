package xmlparser

type ChordConvert struct {
	chordType string
	degree    *Degree
}

func (cc *ChordConvert) convert() string {

	switch cc.chordType {
	case "major":
		return "0"
	case "minor":
		return "1"
	case "diminished", "diminished-seventh":
		return "2"
	case "major-seventh", "major-ninth":
		return "3"
	case "minor-seventh", "minor-ninth":
		if cc.degree != nil {
			if cc.degree.DegreeValue == "5" && cc.degree.DegreeAlter == "-1" {
				return "6"
			}
		}

		return "4"
	case "dominant", "dominant-ninth":

		return "5"
	}

	panic("could not match chord type in ChordConvert:" + cc.chordType)
}
