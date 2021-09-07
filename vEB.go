package vEB

import (
	"fmt"
	"math"
)

// VEB implements a van Embde Boas tree data structure, which efficiently can
// Insert, Delete, Member and find a Succ (successor) all in O(log log u) time.
// Min and Max run in constant time.
type VEB struct {
	min, max int
	local []*VEB
	global []bool
	u, m, q int	//universe size
	swap int	//helper
}

func newVEB(uSize int) *VEB {
	veb := &VEB{}
	veb.u = uSize
	veb.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(veb.u)))))
	veb.q = 1 << int(math.Ceil(math.Log2(math.Sqrt(float64(veb.u)))))
	veb.min, veb.max = -1, -1
	veb.local = make([]*VEB, veb.m)
	veb.global = make([]bool, veb.m)
	return veb
}

// InitVEB inits a new van Emde Boas tree with universe size uSize.
// uSize limits the maximum allowed whole numbers which are insertable to [0, uSize)
// currently, negative numbers are not supported
func InitVEB(uSize int) PrioQ {
	//calculate the smallest larger power of 2 to the given universe size
	u := int(math.Exp2(math.Ceil(math.Log2(float64(uSize)))))
	var v PrioQ = newVEB(u)
	return v
}

// todo: fast Pred/Succ/Member
// todo: sparse mode (hashmaps instead of arrays)

// Insert does not run in O(log log u) time due to lazy initialisation of the tree structure.
// todo: preload mode as init parameter to enable fast inserting.
func (v *VEB) Insert(x int) {
	if v.min == -1 {
		v.min = x
		v.max = x
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
			v.local[i] = newVEB(v.q)
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
func (v *VEB) Delete(x int) {
	//deleting the only element left
	if v.min == v.max {
		v.min = -1
		v.max = -1
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
		//i := v.high(v.min)
		x = i * v.q + v.local[i].min
		v.min = x
	}

	v.local[v.high(x)].Delete(v.low(x))

	if v.local[v.high(x)].min == -1 {
		v.global[v.high(x)] = false
		if x == v.max {

			//find highest remaining entry in v.global
			l := -1
			for k := len(v.global)-1; k >= 0 ; k-- {
				if v.global[k] {
					l = k
					break
				}
			}

			if l == -1 {
				v.max = v.min
			} else {
				v.max = l * v.q + v.local[l].max
			}
		}
	} else {
		if x == v.max {
			v.max = v.high(x)*v.q + v.local[v.high(x)].max
		}
	}

}
func (v *VEB) Succ(x int) int {

	//rekursion end
	if v.u == 2 {
		if x == 0 && v.max == 1 {
			return 1
		} else {
			return 0
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
	return i * v.q + v.local[i].min
}

// not consumed by interface yet, but functional:

func (v *VEB) Member(x int) bool {
	return v.Succ(x-1) == x
}
func (v *VEB) Min(x int) int {
	return v.min
}
func (v *VEB) Max(x int) int {
	return v.max
}
/*
Todo: Predecessor
func (v *VEB) Pred(x int) int {
	return -1
}
*/

// helpers for getting the high and low bits of a number,
// corresponding with its global cluster nummer and its local position in that cluster

func (v *VEB) low(x int) int {
	//todo: replace with bitmasks for speed
	return x % v.q
}
func (v *VEB) high(x int) int {
	//todo: replace with bitmasks for speed
	return x / v.q
}