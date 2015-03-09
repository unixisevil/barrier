package main

import (
	. "barrier"
	"fmt"
	"sync"
)

type thrArg struct {
	incr int
	arr  [narr]int
}

var (
	thrs    [nthr]thrArg
	wg      sync.WaitGroup
	barrier = NewBarrier(nthr)
)

const (
	outloops = 10
	inloops  = 1000
	nthr     = 5
	narr     = 6
)

func thrFunc(arg *thrArg) {
	defer wg.Done()
	for i := 0; i < outloops; i++ {
		barrier.Wait()
		for j := 0; j < inloops; j++ {
			for k := 0; k < narr; k++ {
				arg.arr[k] += arg.incr
			}
		}
		if barrier.Wait() {
			for i := 0; i < nthr; i++ {
				thrs[i].incr += 1
			}
		}
	}
}

func main() {

	for i := 0; i < nthr; i++ {
		thrs[i].incr = i
		for j := 0; j < narr; j++ {
			thrs[i].arr[j] = j + 1
		}
		wg.Add(1)
		go thrFunc(&thrs[i])
	}
	wg.Wait()
	for i := 0; i < nthr; i++ {
		fmt.Printf("%02d: (%d) ", i, thrs[i].incr)
		for j := 0; j < narr; j++ {
			fmt.Printf("%010d ", thrs[i].arr[j])
		}
		fmt.Println()
	}
}
