package doku2md

import (
	"regexp"
)

type Converter struct {
}

func (c Converter) DokuToMd(input string) (output string) {
	line := c.convertHeader(input)
	line = c.convertItalic(line)
	line = c.convertMonospaced(line)
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

	reg := regexp.MustCompile(`(http(s*):)//`)
	if reg.MatchString(output) {
		output = reg.ReplaceAllString(output, `__${1}_placeholder__`)
	}

	reg = regexp.MustCompile(`//(.*)//`)
	if reg.MatchString(output) {
		output = reg.ReplaceAllString(output, "*$1*")
	}

	reg = regexp.MustCompile(`__(http(s*):)_placeholder__`)
	if reg.MatchString(output) {
		output = reg.ReplaceAllString(output, "${1}//")
	}

	return output
}

func (c Converter) convertMonospaced(line string) (output string) {
	output = line
	reg := regexp.MustCompile(`''%%(.*)%%''`)
	if reg.MatchString(output) {
		output = reg.ReplaceAllString(output, "`$1`")
	}
	return output
}
