package internal

import (
	"regexp"
	"sync"

	"github.com/ax4w/gg/internal/pool"
	"github.com/ax4w/gg/internal/search"
)

type (
	GG struct {
		Name   string
		Pool   pool.Pool
		Rx, Ic bool
		P, Q   string
		Cz     int
	}
	result struct {
		sync.Mutex
		result map[int][]string
	}
)

func (g GG) Start() {
	var (
		wg  sync.WaitGroup
		res = result{
			result: make(map[int][]string),
		}
	)
	println("Result for name:", g.Name)
	if g.Ic {
		g.Q = "(?i)" + g.Q
	}
	println("using expr", g.Q)
	expr, err := regexp.Compile(g.Q)
	if err != nil {
		panic(err.Error())
	}
	for i, v := range g.Pool.Lines {
		wg.Add(1)
		go func(lines []string, query string, i int) {
			r := search.Regex(lines, expr, g.Ic)
			res.Lock()
			res.result[i] = r
			res.Unlock()
			wg.Done()
		}(v, g.Q, i)
	}
	wg.Wait()
	for i := 0; i < len(g.Pool.Lines); i++ {
		for _, v := range res.result[i] {
			println(v)
		}
	}
}
