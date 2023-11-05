package vEB

import (
	"math"
)

type V1 struct {
	min, max int
	local    []*V1
	global   *V1
	u, m, q  int //universe size
}

func (v *V1) Init(u int, fullInit bool) {
	if u <= 1 {
		panic("uSize has to be a positive integer larger than 1")
	}

	//calculate the smallest larger power of 2 to the given universe size
	u = int(math.Exp2(math.Ceil(math.Log2(float64(u)))))
	v.init(u, fullInit)
}

func (v *V1) init(uSize int, fullInit bool) {
	v.u = uSize
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = 1 << int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
	v.min, v.max = -1, -1
	v.local = make([]*V1, v.q)

	if fullInit && v.u > 2 {
		for i := range v.local {
			v.local[i] = new(V1)
			v.local[i].init(v.q, fullInit)
		}
	}
}

func (v *V1) Insert(x int) {
	if v.min == -1 {
		v.min, v.max = x, x
		return
	}

	if x < v.min {
		v.min, x = x, v.min
	}

	if v.u > 2 {
		i := high(x, v.q)
		//lazy init
		if v.local[i] == nil {
			v.local[i] = new(V1)
			v.local[i].init(v.q, false)
		}
		low := low(x, v.q)
		if v.local[i].min == -1 {
			//todo: is always insert necessary?
			if v.global == nil {
				v.global = new(V1)
				v.global.init(v.q, false)
				v.global.min = i
				v.global.max = i
			} else {
				v.global.Insert(i)
			}
			v.local[i].min = low
			v.local[i].max = low
		} else {
			v.local[i].Insert(low)
		}
	}

	if x > v.max {
		v.max = x
	}
}

func (v *V1) Succ(x int) int {
	//rekursion end
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		} else {
			//todo: check if at recursive call a detection is necessary
			return -1
		}
	}

	if v.min != -1 && x < v.min {
		return v.min
	}

	//if x < local max, search in local substructure recursively
	l := -1
	hi, lo := high(x, v.q), low(x, v.q)
	if v.global != nil && v.global.Succ(hi-1) == hi {
		l = v.local[hi].max
	}

	if l != -1 && lo < l {
		return hi*v.q + v.local[hi].Succ(lo)
	}

	//global.succ
	i := -1
	if v.global != nil {
		i = v.global.Succ(hi)
	}

	//if no global successor exists, x is bigger than every member
	if i == -1 {
		return -1
	}

	//if global successor exists, return its min
	return i*v.q + v.local[i].min
}

func (v *V1) Pred(x int) int {
	if x == -1 {
		return v.max
	}

	if v.u == 2 {
		if x == 1 && v.min == 0 {
			return 0
		}
		return -1
	}
	if v.global == nil {
		if v.min < x {
			return v.min
		}
		return -1
	}

	xHi, xLo := high(x, v.q), low(x, v.q)

	if v.global.Member(xHi) {
		if v.local[xHi].min < x && v.local[xHi].min > -1 {
			return int(xHi)*v.q + v.local[xHi].Pred(xLo)
		}
	}

	//global.pred
	gloPred := v.global.Pred(xHi)

	if gloPred >= 0 {
		return gloPred*v.q + v.local[gloPred].max
	}
	if v.min >= 0 && v.min < x {
		return v.min
	}
	return -1
}

func (v *V1) Delete(x int) {
	//deleting the only element left
	if v.min == v.max {
		//only delete, if element is present in tree
		if v.min == x {
			v.min = -1
			v.max = -1
		}
		return
	}

	//deleting the second last element
	if v.u == 2 {
		v.min = 1 - x
		v.max = v.min
		return
	}

	//delete second smallest element from substructure and overwrite min (where x was stored)
	if x == v.min {
		i := v.global.min

		x = i*v.q + v.local[i].min
		v.min = x
	}
	hi, lo := high(x, v.q), low(x, v.q)
	v.local[hi].Delete(lo)

	if v.local[hi].min == -1 {
		v.global.Delete(hi)
		if x == v.max {

			//find highest remaining entry in v.global
			l := v.Pred(x)

			if l == -1 {
				v.max = v.min
			} else {
				v.max = l*v.q + v.local[l].max
			}
		}
	} else {
		if x == v.max {
			v.max = hi*v.q + v.local[hi].max
		}
	}
}

func (v *V1) Member(x int) bool {
	return v.Succ(x-1) == x
}
func (v *V1) Min() int {
	return v.min
}
func (v *V1) Max() int {
	return v.max
}

func low(x, b int) int {
	//todo: replace with bitmasks for speed
	return x % b
}
func high(x, b int) int {
	//todo: replace with bitmasks for speed
	return x / b
}
