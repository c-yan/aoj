package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

const (
	defaultValue = math.MaxInt64
	defaultLazy  = 0
)

type segmentTree struct {
	offset int
	values []int
	lazy   []int
}

func newSegmentTree(n int) segmentTree {
	var result segmentTree
	t := 1
	for t < n {
		t *= 2
	}
	result.offset = t - 1
	result.values = make([]int, 2*t-1)
	result.lazy = make([]int, 2*t-1)
	for i := 0; i < 2*t-1; i++ {
		result.values[i] = defaultValue
		result.lazy[i] = defaultLazy
	}
	return result
}

func op(x, y int) int {
	return min(x, y)
}

func (st segmentTree) build(a []int) {
	for i, v := range a {
		st.values[st.offset+i] = v
	}
	for i := st.offset - 1; i > -1; i-- {
		st.values[i] = op(st.values[i*2+1], st.values[i*2+2])
	}
}

func (st segmentTree) propagateTo(index int, value int) {
	st.lazy[index] += value
	st.values[index] += value
}

func (st segmentTree) propagateAt(index int) {
	if index != 0 {
		st.propagateAt(st.getParentIndex(index))
	}
	if st.lazy[index] == defaultLazy {
		return
	}
	st.propagateTo(index*2+1, st.lazy[index])
	st.propagateTo(index*2+2, st.lazy[index])
	st.lazy[index] = defaultLazy
}

func (st segmentTree) recalcAt(index int) {
	st.values[index] = op(st.values[index*2+1], st.values[index*2+2])
	if index != 0 {
		st.recalcAt(st.getParentIndex(index))
	}
}

func (st segmentTree) applyAt(index int, value int) {
	st.lazy[index] += value
	st.values[index] += value
	if index != 0 {
		st.recalcAt(st.getParentIndex(index))
	}
}

func (st segmentTree) getParentIndex(index int) int {
	if index < 1 {
		panic("BUG: doesn't exists parent of 0")
	}
	return (index - 1) / 2
}

func (st segmentTree) propagate(start, stop int) {
	l := start + st.offset
	r := stop + st.offset
	for l < r {
		if l&1 == 0 {
			if l != 0 {
				st.propagateAt(st.getParentIndex(l))
			}
		}
		if r&1 == 0 {
			if r-1 != 0 {
				st.propagateAt(st.getParentIndex(r - 1))
			}
		}
		l = l / 2
		r = (r - 1) / 2
	}
}

func (st segmentTree) apply(start, stop int, value int) {
	st.propagate(start, stop)
	l := start + st.offset
	r := stop + st.offset
	for l < r {
		if l&1 == 0 {
			st.applyAt(l, value)
		}
		if r&1 == 0 {
			st.applyAt(r-1, value)
		}
		l = l / 2
		r = (r - 1) / 2
	}
}

func (st segmentTree) query(start, stop int) int {
	st.propagate(start, stop)
	result := defaultValue
	l := start + st.offset
	r := stop + st.offset
	for l < r {
		if l&1 == 0 {
			result = op(result, st.values[l])
		}
		if r&1 == 0 {
			result = op(result, st.values[r-1])
		}
		l = l / 2
		r = (r - 1) / 2
	}
	return result
}

func main() {
	defer flush()

	n := readInt()
	q := readInt()

	st := newSegmentTree(n)
	st.build(make([]int, n))
	for i := 0; i < q; i++ {
		a := readInt()
		if a == 0 {
			s := readInt()
			t := readInt()
			x := readInt()
			st.apply(s, t+1, x)
		} else if a == 1 {
			s := readInt()
			t := readInt()
			println(st.query(s, t+1))
		}
	}
}

const (
	ioBufferSize = 1 * 1024 * 1024 // 1 MB
)

var stdinScanner = func() *bufio.Scanner {
	result := bufio.NewScanner(os.Stdin)
	result.Buffer(make([]byte, ioBufferSize), ioBufferSize)
	result.Split(bufio.ScanWords)
	return result
}()

func readString() string {
	stdinScanner.Scan()
	return stdinScanner.Text()
}

func readInt() int {
	result, err := strconv.Atoi(readString())
	if err != nil {
		panic(err)
	}
	return result
}

var stdoutWriter = bufio.NewWriter(os.Stdout)

func flush() {
	stdoutWriter.Flush()
}

func println(args ...interface{}) (int, error) {
	return fmt.Fprintln(stdoutWriter, args...)
}
