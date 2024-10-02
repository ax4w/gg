package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/ax4w/gg/internal"
	"github.com/ax4w/gg/internal/pool"
)

var (
	//regex      = flag.Bool("rx", false, "search using an regular expression")
	recursive  = flag.Bool("rec", false, "traverse sub folders in a dir")
	ignoreCase = flag.Bool("ic", false, "ignore upper and lower casing")
	path       = flag.String("p", "", "the path to the folder / file")
	query      = flag.String("q", "", "the string to search for")
	chunkSize  = flag.Int("cz", 15000, "chunk size per thread")
	pipe       string
)

func init() {
	flag.Parse()
	if len(*query) == 0 {
		panic("No query was provided")
	}
	if len(*path) == 0 {
		panic("No path or piped in text was provided")
	}
	if *chunkSize == 0 {
		panic("Chunk Size can't be 0")
	}
}

func file(p string) {
	bytes, err := os.ReadFile(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		return
	}
	internal.GG{
		Name: p,
		Pool: pool.New(strings.Split(string(bytes), "\n"), *chunkSize),
		Rx:   false,
		Ic:   *ignoreCase,
		P:    p,
		Q:    regexp.QuoteMeta(strings.TrimSpace(*query)),
		Cz:   *chunkSize,
	}.Start()
}

func folder(p string) {
	entries, err := os.ReadDir(p)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		return
	}
	var wg sync.WaitGroup
	for _, entry := range entries {
		if entry.IsDir() {
			if *recursive {
				wg.Add(1)
				go func(p, name string) {
					folder(filepath.Join(p, name))
					wg.Done()
				}(p, entry.Name())
			}
		} else {
			wg.Add(1)
			go func(p, name string) {
				file(filepath.Join(p, name))
				wg.Done()
			}(p, entry.Name())

		}
	}
	wg.Wait()
}

func piped() {
	//lines := strings.Split(pipe, "\n")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(pipe) > 0 {
		piped()
	} else {
		f, err := os.Stat(*path)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			return
		}
		if f.IsDir() {
			folder(*path)
		} else {
			file(*path)
		}

	}

}
