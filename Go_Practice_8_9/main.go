package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//var vFlag = flag.Bool("v", false, "show verbose progress messages")

type SizeResponse struct {
	root int
	size int64
}

func main() {

	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	sizeResponses := make(chan SizeResponse)
	var n sync.WaitGroup
	for i, root := range roots {
		n.Add(1)
		go walkDir(root, &n, i, sizeResponses)
	}
	go func() {
		n.Wait()
		close(sizeResponses)
	}()

	var tick <-chan time.Time
	// if *vFlag {
	// 	tick = time.Tick(500 * time.Millisecond)
	// }
	nfiles := make([]int64, len(roots))
	nbytes := make([]int64, len(roots))

loop:
	for {
		select {
		case sr, ok := <-sizeResponses:
			if !ok {
				break loop
			}
			nfiles[sr.root]++
			nbytes[sr.root] += sr.size
		case <-tick:
			printDiskUsage(roots, nfiles, nbytes)
		}
	}

	printDiskUsage(roots, nfiles, nbytes)

}

func printDiskUsage(roots []string, nfiles, nbytes []int64) {
	for i, r := range roots {
		fmt.Printf("%10d files  %.3f GB under %s\n", nfiles[i], float64(nbytes[i])/1e9, r)
	}
}

func walkDir(dir string, n *sync.WaitGroup, root int, sizeResponses chan<- SizeResponse) {
	defer n.Done()
	file, err := os.Open(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return
	}
	defer file.Close()

	entries, err := file.Readdir(0)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, root, sizeResponses)
		} else {
			sizeResponses <- SizeResponse{root, entry.Size()}
		}
	}
}
