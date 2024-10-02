package search

import (
	"regexp"
	"sort"
)

const (
	colorStart = "\033[93m\033[1m"
	colorEnd   = "\033[0m"
)

func color(str string, locations [][]int) string {
	if len(locations) == 0 {
		return str
	}

	sort.Slice(locations, func(i, j int) bool {
		return locations[i][0] > locations[j][0]
	})

	result := str
	for _, loc := range locations {
		start, end := loc[0], loc[1]
		result = result[:start] + colorStart + result[start:end] + colorEnd + result[end:]
	}

	return result
}
func Regex(lines []string, query *regexp.Regexp, ic bool) (res []string) {
	for _, v := range lines {
		if locs := query.FindAllIndex([]byte(v), -1); locs != nil {
			res = append(res, color(v, locs))
		}
	}
	return res
}
