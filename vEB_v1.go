package vEB

import "math"

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
	v.local = make([]*V1, v.m)
	v.global = new(V1)
	v.global.init(v.q, fullInit)

	if fullInit && v.u > 2 {
		for i := range v.local {
			v.local[i] = new(V1)
			v.local[i].init(v.q, fullInit)
		}
	}
}

func (v *V1) Insert(x int) {

}

func (v *V1) Succ(x int) int {

	return -1
}

func (v *V1) Delete(x int) {

}
