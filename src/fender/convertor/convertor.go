package convertor

import (
	"fender/utility"
	"regexp"
)

type direction int

const (
	endOfDos  = "\r\n"
	endOfUnix = "\n"

	regexPatternDos2Unix = "\r\n|\n|\r/\r\n"
	regexPatternUnix2Dos = "\r\n|\n|\r/\n"

	directionDos2Unix direction = 0
	directionUnix2Dos direction = 1
)

func lineEndConvertor(line []byte, d direction) []byte {
	var regexPattern string
	var replaceWord string

	if d == directionDos2Unix {
		regexPattern = regexPatternDos2Unix
		replaceWord = endOfUnix
	} else {
		regexPattern = regexPatternUnix2Dos
		replaceWord = endOfDos
	}

	r, _ := regexp.Compile(regexPattern)
	line = r.ReplaceAll(line, utility.Str2bytes(replaceWord))

	return line
}

// Conv2Dos will convert all unix line end to dos line end
func Conv2Dos(line []byte) []byte {
	return lineEndConvertor(line, directionUnix2Dos)
}

// Conv2Unix will convert all dos line end to unix line end
func Conv2Unix(line []byte) []byte {
	return lineEndConvertor(line, directionDos2Unix)
}
