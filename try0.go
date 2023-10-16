package vEB

import (
	"fmt"
	"math"
)

type Try0 struct {
	global []bool
	local  [][]int
	q      int
}

func (v *Try0) Init(u int, fullInit bool) {
	//v.u = int(math.Exp2(math.Ceil(math.Log2(float64(u)))))
	//v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = 1 << int(math.Ceil(math.Log2(math.Sqrt(float64(u)))))
	v.global = make([]bool, v.q)
	v.local = make([][]int, v.q)
}

func (v *Try0) Insert(x int) {
	xHi := x / v.q
	if !v.global[xHi] {
		v.global[xHi] = true
		v.local[xHi] = make([]int, 0)
	}

	loc := v.local[xHi]

	small := -1
	for i, val := range loc {
		if val == x {
			return
		}
		if val < x {
			small = i
		} else {
			break
		}
	}

	loc = append(loc, 0)
	if len(loc)-small > 2 {
		copy(loc[small+2:], loc[small+1:])
	}
	loc[small+1] = x
	v.local[xHi] = loc
}
func (v *Try0) Succ(x int) int {
	xHi := x / v.q

	if v.global[xHi] {
		for _, val := range v.local[xHi] {
			if x < val {
				return val
			}
		}
	}

	if xHi == v.q-1 {
		// end of global
		return -1
	}
	toSearch := -1
	for i := xHi + 1; i < v.q; i++ {
		if v.global[i] {
			toSearch = i
			break
		}
	}
	if toSearch == -1 {
		return -1
	}

	for _, val := range v.local[toSearch] {
		if x < val {
			return val
		}
	}

	fmt.Println("This should never print... :^  )")
	return -1
}
func (v *Try0) Pred(x int) int {
	xHi := x / v.q

	if v.global[xHi] {
		loc := v.local[xHi]
		for i := len(loc) - 1; i >= 0; i-- {
			if x > loc[i] {
				return loc[i]
			}
		}
	}

	if xHi == 0 {
		// end of global
		return -1
	}
	toSearch := -1
	for i := xHi - 1; i >= 0; i-- {
		if v.global[i] {
			toSearch = i
			break
		}
	}
	if toSearch == -1 {
		return -1
	}

	loc := v.local[toSearch]
	for i := len(loc) - 1; i >= 0; i-- {
		if x > loc[i] {
			return loc[i]
		}
	}

	fmt.Println("This should never print... :^  )")
	return -1
}
func (v *Try0) Delete(x int) {
	xHi := x / v.q
	loc := v.local[xHi]
	if !v.global[xHi] {
		return
	}
	if len(loc) == 1 && x == loc[0] {
		v.global[xHi] = false
		v.local[xHi] = nil
		return
	}

	ind := -1
	for i, val := range loc {
		if x < val {
			// early abort, as local is sorted
			break
		}
		if val == x {
			ind = i
			break
		}
	}
	if ind == -1 {
		return
	}
	loc = append(loc[:ind], loc[ind+1:]...)
	v.local[xHi] = loc
}
func (v *Try0) Min() int {
	xHi := -1
	for i, b := range v.global {
		if b {
			xHi = i
			break
		}
	}
	if xHi == -1 {
		return -1
	}

	return v.local[xHi][0]
}
func (v *Try0) Max() int {
	xHi := -1
	for i := v.q - 1; i >= 0; i-- {
		if v.global[i] {
			xHi = i
			break
		}
	}
	if xHi == -1 {
		return -1
	}

	return v.local[xHi][len(v.local[xHi])-1]
}
func (v *Try0) Member(x int) bool {
	xHi := x / v.q
	if !v.global[xHi] {
		return false
	}

	loc := v.local[xHi]
	for _, val := range loc {
		if x < val {
			// early abort due to loc being sorted
			return false
		}
		if x == val {
			return true
		}
	}
	return false
}
