package doku2md

import (
	"regexp"
)

type Converter struct {
}

func (c Converter) DokuToMd(input string) (output string) {
	line := c.convertHeader(input)
	line = c.convertItalic(line)
	return line
}

func (c Converter) convertHeader(line string) (output string) {
	output = line
	patterns := []string{`====== (.*) ======`, `===== (.*) =====`, `==== (.*) ====`, `=== (.*) ===`, `== (.*) ==`, `= (.*) =`}
	repls := []string{"# $1", "## $1", "### $1", "#### $1", "##### $1", "###### $1"}

	for i, pattern := range patterns {
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(output) {
			output = reg.ReplaceAllString(output, repls[i])
			break
		}
	}

	return output
}

func (c Converter) convertItalic(line string) (output string) {
	output = line
	reg := regexp.MustCompile(`//(.*)//`)
	if reg.MatchString(output) {
		output = reg.ReplaceAllString(output, "*$1*")
	}
	return output
}
