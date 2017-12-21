package xmlparser

import (
	"strings"
	"fmt"
	"errors"
	"strconv"
)

type NoteDuration struct {
	Type string
	Dotted bool
}

type Validate struct {
	Bars []Bar
}

func (v *Validate)CheckDurations() error {
	for i, bar := range v.Bars {
		if i == 0 {
			continue
		}

		isValid, _ := isBarDurationValid(bar)

		//fmt.Printf("Bar #%s has Valid status : %s, with duration: %s \n",strconv.Itoa(i), strconv.FormatBool(isValid), strconv.Itoa(totalDuration))

		if !isValid {
			return errors.New("BAR #" + strconv.Itoa(i) + " has INVALID durations!")
		}
	}

	return nil
}

func durationsFromBar(bar Bar)[]NoteDuration {

	var noteDurations []NoteDuration

	for _, note := range bar.Notes {
		dotted := strings.Contains(fmt.Sprintf("%v", note.Dot), "dot")
		noteDurations = append(noteDurations, NoteDuration{
			Type:   note.Type,
			Dotted: dotted,
		})
	}

	return noteDurations
}

func isBarDurationValid(bar Bar) (bool, int) {

	noteDurations := durationsFromBar(bar)
	total := 0

	for _, duration := range noteDurations {
		total = total + getDuration(duration.Type, duration.Dotted)
	}

	return total == 16, total
}

func getDuration(durType string, dotted bool) int {

	if dotted == false {
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
			//not supported
		} else if durType == "eighth" {
			return 3
		} else if durType == "quarter" {
			return 6
		} else if durType == "half" {
			return 12
		}
	}

	panic("UNSUPPORTED NOT DURATION!! " + durType + " DOTTED: " + strconv.FormatBool(dotted))
}
