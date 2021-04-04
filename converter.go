package doku2md

import (
	"regexp"
)

type Converter struct {
}

func (c *Converter) DokuToMd(input string) (output string) {
	patterns := []string{`====== (.*) ======`, `===== (.*) =====`, `==== (.*) ====`, `=== (.*) ===`, `== (.*) ==`, `= (.*) =`}
	repls := []string{"# $1", "## $1", "### $1", "#### $1", "##### $1", "###### $1"}

	for i, pattern := range patterns {
		reg := regexp.MustCompile(pattern)
		if reg.MatchString(input) {
			output = reg.ReplaceAllString(input, repls[i])
			break
		}
	}

	return output
}
