package internal

import (
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
		result []string
	}
)

func (g GG) Start() {
	var wg sync.WaitGroup
	println("Result for name:", g.Name)
	for _, v := range g.Pool.Lines {
		wg.Add(1)
		go func(lines []string, query string) {
			for _, v := range search.String(lines, query, g.Ic) {
				println(v)
				//println("\033[93m \033[1m", v, "\033[0m")
			}
			wg.Done()
		}(v, g.Q)
	}
	wg.Wait()
}
