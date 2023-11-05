package vEB

import (
	"fmt"
	"math"
)

type Try4 struct {
	global     *Try4
	local      []*Try4
	globalLeaf *BitsPrioQ
	localLeaf  []*BitsPrioQ
	u, q, m    int
	q_         int //log2(q)
	min, max   int
}

const BIT_CUTOFF = 64

func (v *Try4) Init(u int, fullInit bool) {
	v.u = 1 << int(math.Ceil(math.Log2(float64(u))))
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q_ = int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = v.u / v.m // == u/m
	v.min, v.max = -1, -1

	if u > BIT_CUTOFF { // todo: test where cutoff?
		v.global = &Try4{}
		v.global.Init(v.m, fullInit)

		v.local = make([]*Try4, v.m)
		for i, _ := range v.local {
			a := &Try4{}
			a.Init(v.q, fullInit)
			v.local[i] = a
		}
	} else if u > 2 {
		v.globalLeaf = &BitsPrioQ{}
		v.globalLeaf.Init(v.m, fullInit)

		v.localLeaf = make([]*BitsPrioQ, v.m)
		for i := range v.localLeaf {
			a := &BitsPrioQ{}
			a.Init(v.q, fullInit)
			v.localLeaf[i] = a
		}
	}
}
func (v *Try4) Insert(x int) {
	if v.min < 0 {
		v.min, v.max = x, x
		return
	}
	if x < v.min {
		// swap and continue
		v.min, x = x, v.min
	}
	if x > v.max {
		v.max = x
	}
	if v.u > 2 {
		xHi := v.high(x)
		if v.local[xHi].min < 0 {
			v.global.Insert(int(xHi))
		}
		v.local[xHi].Insert(int(v.low(x)))
	}
}
func (v *Try4) Succ(x int) int {
	if x < 0 {
		return v.min
	}

	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		}

		return -1
	}
	if v.min >= 0 && x < v.min {
		return v.min
	}

	xHi, xLo := v.high(x), v.low(x)
	maxLo := (v.local[xHi]).max
	if maxLo >= 0 && int(xLo) < maxLo {
		return int(xHi)*v.q + (v.local[xHi]).Succ(int(xLo))
	}

	gloSucc := v.global.Succ(int(xHi))
	if gloSucc < 0 {
		return -1
	}
	return gloSucc*v.q + (v.local[gloSucc]).min
}
func (v *Try4) Pred(x int) int {
	// todo:
	return -1
}
func (v *Try4) Delete(x int) {
	if x < 0 {
		return
	}
	if v.max == v.min {
		// guard delete wrong number
		if v.min != x {
			return
		}
		v.max, v.min = -1, -1
		return
	}
	if v.u == 2 {
		v.min = 1 - x
		v.max = v.min
		return
	}
	if v.min == x {
		gMin := v.global.min
		x = gMin*v.q + (v.local[gMin]).min
		v.min = x
	}
	xHi, xLo := v.high(x), v.low(x)
	if xHi > 100000000 {
		fmt.Println("lol")
	}
	(v.local[xHi]).Delete(int(xLo))
	if (v.local[xHi]).min < 0 {
		v.global.Delete(int(xHi))
		if x == v.max {
			gloMax := v.global.max
			if gloMax < 0 {
				v.max = v.min
			} else {
				v.max = gloMax*v.q + (v.local[gloMax]).max
			}
		}
	} else {
		if x == v.max {
			v.max = int(xHi)*v.q + (v.local[int(xHi)]).max
		}
	}
}
func (v *Try4) Min() int {
	return v.min
}
func (v *Try4) Max() int {
	return v.max
}
func (v *Try4) Member(x int) bool {
	return v.Succ(x-1) == x
}

func (v *Try4) high(x int) uint64 {
	return uint64(x) >> v.q_
}
func (v *Try4) low(x int) uint64 {
	return uint64(x) & (^uint64(0) >> (64 - v.q_))
}
