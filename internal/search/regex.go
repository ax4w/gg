package search

import (
	"regexp"
	"strings"
)

func Regex(lines []string, query *regexp.Regexp, ic bool) (res []string) {
	for _, v := range lines {
		if ic {
			if loc := query.FindStringIndex(strings.ToLower(v)); loc != nil {
				res = append(res, v)
			}
		} else {
			if loc := query.FindStringIndex(v); loc != nil {
				res = append(res, v)
			}
		}
	}
	return res
}
