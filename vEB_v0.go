package vEB

import (
	"fmt"
	"math"
)

// V0 is the first Iteration to implement a van Embde Boas tree data structure, which efficiently can
// Insert, Delete, Member and find a Succ (successor) all in O(log log u) time.
// Min and Max run in constant time.
// This Implementation is not a full vEB as it doesn't have recursive globals and the fullInit is wrong.
type V0 struct {
	min, max int
	local    []*V0
	global   []bool
	u, m, q  int //universe size
	swap     int //helper
}

func (v *V0) init(uSize int, fullInit bool) {
	v.u = uSize
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = 1 << int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
	v.min, v.max = -1, -1
	v.local = make([]*V0, v.m)
	v.global = make([]bool, v.m)

	if fullInit && v.u > 2 {
		for i := range v.local {
			v.local[i] = new(V0)
			v.local[i].init(v.q, fullInit)
		}
	}
}

// Init inits a new Tree with universe size uSize.
// uSize limits the maximum allowed whole numbers which are insertable to [0, uSize)
func (v *V0) Init(u int, fullInit bool) {
	if u <= 1 {
		panic("uSize has to be a positive integer larger than 1")
	}

	//calculate the smallest larger power of 2 to the given universe size
	u = int(math.Exp2(math.Ceil(math.Log2(float64(u)))))
	v.init(u, fullInit)
}

// Insert does not run in O(log log u) time if lazy initialisation of the tree structure is turned on.
func (v *V0) Insert(x int) {
	if v.min == -1 {
		v.min, v.max = x, x
		return
	}

	if x < v.min {
		v.swap = v.min
		v.min = x
		x = v.swap
	}

	if v.u > 2 {
		i := v.high(x)
		//lazy init
		if v.local[i] == nil {
			v.local[i] = new(V0)
			v.local[i].init(v.q, false)
		}
		if v.local[i].min == -1 {
			v.global[i] = true
			v.local[i].min = v.low(x)
			v.local[i].max = v.low(x)
		} else {
			v.local[i].Insert(v.low(x))
		}
	}

	if x > v.max {
		v.max = x
	}
}
func (v *V0) Delete(x int) {
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
		i := -1
		//should always find at least one
		for j := 0; j < len(v.global); j++ {
			if v.global[j] {
				i = j
				break
			}
		}
		if i == -1 {
			fmt.Println("Should never happen, global min not found...")
		}

		x = i*v.q + v.local[i].min
		v.min = x
	}

	v.local[v.high(x)].Delete(v.low(x))

	if v.local[v.high(x)].min == -1 {
		v.global[v.high(x)] = false
		if x == v.max {

			//find highest remaining entry in v.global
			l := -1
			for k := len(v.global) - 1; k >= 0; k-- {
				if v.global[k] {
					l = k
					break
				}
			}

			if l == -1 {
				v.max = v.min
			} else {
				v.max = l*v.q + v.local[l].max
			}
		}
	} else {
		if x == v.max {
			v.max = v.high(x)*v.q + v.local[v.high(x)].max
		}
	}

}
func (v *V0) Succ(x int) int {

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
	if v.global[v.high(x)] {
		l = v.local[v.high(x)].max
	}

	if l != -1 && v.low(x) < l {
		return v.high(x)*v.q + v.local[v.high(x)].Succ(v.low(x))
	}

	//global.succ
	i := -1
	for j := v.high(x) + 1; j < len(v.global); j++ {
		if v.global[j] {
			i = j
			break
		}
	}

	//if no global successor exists, x is bigger than every member
	if i == -1 {
		return -1
	}

	//if global successor exists, return its min
	return i*v.q + v.local[i].min
}
func (v *V0) Member(x int) bool {
	return v.Succ(x-1) == x
}
func (v *V0) Min() int {
	return v.min
}
func (v *V0) Max() int {
	return v.max
}

// Todo: Predecessor
func (v *V0) Pred(x int) int {
	if x == -1 {
		return v.max
	}

	if v.u == 2 {
		if x == 1 && v.min == 0 {
			return 0
		}
		return -1
	}
	xHi := v.high(x)
	if v.global[xHi] {
		if v.local[xHi].min < x && v.local[xHi].min > -1 {
			return int(xHi)*v.q + v.local[xHi].Pred(int(v.low(x)))
		}
	}

	//global.pred
	gloPred := -1
	for j := v.high(x) - 1; j >= 0; j-- {
		if v.global[j] {
			gloPred = j
			break
		}
	}

	if gloPred >= 0 {
		return gloPred*v.q + v.local[gloPred].max
	}
	if v.min >= 0 && v.min < x {
		return v.min
	}
	return -1
}

// helpers for getting the high and low bits of a number,
// corresponding with its global cluster nummer and its local position in that cluster
func (v *V0) low(x int) int {
	//todo: replace with bitmasks for speed
	return x % v.q
}
func (v *V0) high(x int) int {
	//todo: replace with bitmasks for speed
	return x / v.q
}
