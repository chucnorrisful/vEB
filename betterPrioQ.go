package vEB

import "math"

type BetterPrioQ struct {
	global  []bool
	local   []BetterPrioQ
	u, q, m int
}

func (v *BetterPrioQ) Init(u int, fullInit bool) {
	v.u = int(math.Exp2(math.Ceil(math.Log2(float64(u)))))
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = 1 << int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
}

func (v *BetterPrioQ) Insert(x int) {

}

func (v *BetterPrioQ) Succ(x int) int {

	return -1
}

func (v *BetterPrioQ) Delete(x int) {

}
