package validate

type NoteDuration struct {
	Type string
	Dotted bool
}

func BarDuration(noteDurations []NoteDuration) bool {

	total := 0

	for _, duration := range noteDurations {
		total = total + getDuration(duration.Type, duration.Dotted)
	}

	return total == 16
}

func getDuration(durType string, dotted bool) int {

	if dotted == true {
		if durType == "16th" {
			return 1
		} else if durType == "eighth" {
			return 2
		} else if durType == "quarter" {
			return 4
		} else if durType == "half" {
			return 8
		} else if durType == "whole" {
			return 16
		}
	}else{
		if durType == "16th" {

		} else if durType == "eighth" {
			return 3
		} else if durType == "quarter" {
			return 6
		} else if durType == "half" {
			return 12
		}
	}

	panic("UNSUPPORTED NOT DURATION!! " + durType)
}
