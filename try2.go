package vEB

import (
	"math"
)

type Try2 struct {
	global   PrioQ
	local    []*PrioQ
	u, q, m  int
	q_       int //log2(q)
	min, max int
}

const MEMBER2 = false

func (v *Try2) Init(u int, fullInit bool) {
	v.u = 1 << int(math.Ceil(math.Log2(float64(u))))
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q_ = int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = v.u / v.m // == u/m
	v.min, v.max = -1, -1

	if u > 2 {
		v.global = &Try2{}
		v.global.Init(v.m, fullInit)

		v.local = make([]*PrioQ, v.m)
		for i, _ := range v.local {
			var a PrioQ = &Try2{}
			a.Init(v.q, fullInit)
			v.local[i] = &a
		}
	}
}
func (v *Try2) Insert(x int) {
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
		xHi, xLo := v.split(x)
		loc := *v.local[xHi]
		if loc.Min() < 0 {
			v.global.Insert(int(xHi))
		}
		loc.Insert(int(xLo))
	}
}
func (v *Try2) Succ(x int) int {
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

	xHi, xLo := v.split(x)
	maxLo := (*v.local[xHi]).Max()
	if maxLo >= 0 && int(xLo) < maxLo {
		return int(xHi)*v.q + (*v.local[xHi]).Succ(int(xLo))
	}

	gloSucc := v.global.Succ(int(xHi))
	if gloSucc < 0 {
		return -1
	}
	return gloSucc*v.q + (*v.local[gloSucc]).Min()
}
func (v *Try2) Pred(x int) int {
	// todo:
	return -1
}
func (v *Try2) Delete(x int) {
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
		gMin := v.global.Min()
		x = gMin*v.q + (*v.local[gMin]).Min()
		v.min = x
	}
	xHi, xLo := v.split(x)
	(*v.local[xHi]).Delete(int(xLo))
	if (*v.local[xHi]).Min() < 0 {
		v.global.Delete(int(xHi))
		if x == v.max {
			gloMax := v.global.Max()
			if gloMax < 0 {
				v.max = v.min
			} else {
				v.max = gloMax*v.q + (*v.local[gloMax]).Max()
			}
		}
	} else {
		if x == v.max {
			v.max = int(xHi)*v.q + (*v.local[int(xHi)]).Max()
		}
	}
}
func (v *Try2) Min() int {
	return v.min
}
func (v *Try2) Max() int {
	return v.max
}
func (v *Try2) Member(x int) bool {
	if MEMBER2 {
		return v.Succ(x-1) == x
	}

	if v.min < 0 {
		return false
	}
	if x < v.min {
		return false
	}
	if x > v.max {
		return false
	}
	if x == v.min {
		return true
	}
	if x == v.max {
		return true
	}
	if v.u == 2 {
		// as above checked, x is eighter min or max, and thus is a member
		return true
	}

	xHi, xLo := v.split(x)

	if !v.global.Member(int(xHi)) {
		return false
	}

	return (*v.local[xHi]).Member(int(xLo))
}

func (v *Try2) split(x int) (hi, lo uint64) {
	x2 := uint64(x)

	hi = x2 >> v.q_
	lo = x2 & (^uint64(0) >> (64 - v.q_))
	return
}
