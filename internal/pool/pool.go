package pool

type Pool struct {
	Lines [][]string
}

func New(lines []string, cz int) Pool {
	var p Pool
	for i := 0; i < len(lines); i += cz {
		var chunk []string
		if i+cz < len(lines) {
			chunk = lines[i : i+cz]
		} else {
			chunk = lines[i : len(lines)-1]
		}
		p.Lines = append(p.Lines, chunk)
	}
	return p
}
