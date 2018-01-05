package xmlparser

type KeyConvert struct {
	Key int
}

func (kc *KeyConvert) convert() string {

	switch kc.Key {
	case 0:
		return "31"
	case -1:
		return "61"
	case -2:
		return "12"
	case -3:
		return "42"
	case -4:
		return "72"
	case -5:
		return "32"
	case 1:
		return "71"
	case 2:
		return "41"
	case 3:
		return "11"
	case 4:
		return "51"
	case 5:
		return "21"
	case 6:
		return "62"
	case 7:
		return "32"
	}

	panic("key not found in KeyConvert")
}
