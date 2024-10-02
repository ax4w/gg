package search

import "strings"

func color(str string, start, ln int) string {
	return str[:start] + "\033[93m \033[1m" + str[start:ln] + "\033[0m" + str[ln:]
}

func String(lines []string, query string, ic bool) (res []string) {
	for _, v := range lines {
		if ic {
			if i := strings.Index(strings.ToLower(v), strings.ToLower(query)); i != -1 {
				res = append(res, color(v, i, i+len(query)))
			}
		} else {
			if i := strings.Index(v, query); i != -1 {
				res = append(res, color(v, i, i+len(query)))
			}
		}
	}
	return res
}
