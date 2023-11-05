package vEB

import (
	"math"
)

type Try1 struct {
	global  PrioQ
	local   []*PrioQ
	u, q, m int
	q_      int //log2(q)
}

// Try1Leaf handles (and stores) the recursion-end case
type Try1Leaf struct {
	a, b int
}

func (t *Try1Leaf) Insert(x int) {
	if x == 0 {
		t.a = 0
	}
	if x == 1 {
		t.b = 1
	}
}
func (t *Try1Leaf) Delete(x int) {
	if x == 0 {
		t.a = -1
	}
	if x == 1 {
		t.b = -1
	}
}
func (t *Try1Leaf) Succ(x int) int {
	if x == 0 {
		return t.b
	}
	return -1
}
func (t *Try1Leaf) Pred(x int) int {
	if x == 1 {
		return t.a
	}
	return -1
}
func (t *Try1Leaf) Member(x int) bool {
	return x == t.a || x == t.b
}
func (t *Try1Leaf) Min() int {
	if t.a != -1 {
		return t.a
	} else if t.b != -1 {
		return t.b
	} else {
		return -1
	}
}
func (t *Try1Leaf) Max() int {
	if t.b != -1 {
		return t.b
	} else if t.a != -1 {
		return t.a
	} else {
		return -1
	}
}
func (t *Try1Leaf) Init(u int, fullInit bool) {
	t.a, t.b = -1, -1
}

func (v *Try1) Init(u int, fullInit bool) {
	if !fullInit {
		//fmt.Println("only full init supported atm. Init runs in O(u).")
	}

	v.u = 1 << int(math.Ceil(math.Log2(float64(u))))
	v.m = 1 << int(math.Floor(math.Log2(math.Sqrt(float64(v.u)))))
	v.q_ = int(math.Ceil(math.Log2(math.Sqrt(float64(v.u)))))
	v.q = v.u / v.m // == u/m

	v.local = make([]*PrioQ, v.m)
	for i, _ := range v.local {
		var a PrioQ
		if v.q > 2 {
			a = &Try1{}
		} else {
			a = &Try1Leaf{}
		}

		v.local[i] = &a
		(*v.local[i]).Init(v.q, fullInit)
	}

	if v.m > 2 {
		v.global = &Try1{}
	} else {
		v.global = &Try1Leaf{}
	}

	v.global.Init(v.m, fullInit)
}
func (v *Try1) Insert(x int) {
	if x < 0 {
		return
	}
	xHi, xLo := v.split(x)
	v.global.Insert(int(xHi))
	(*v.local[xHi]).Insert(int(xLo))
}
func (v *Try1) Succ(x int) int {
	if x < 0 {
		gMin := v.global.Min()
		if gMin == -1 {
			return -1
		}
		return gMin*v.q + (*v.local[gMin]).Min()
	}

	xHi, xLo := v.split(x)
	loSucc := (*v.local[xHi]).Succ(int(xLo))
	if loSucc != -1 {
		return int(xHi)*v.q + loSucc
	}

	hiSucc := v.global.Succ(int(xHi))
	if hiSucc == -1 {
		return -1
	}

	return int(hiSucc)*v.q + (*v.local[hiSucc]).Min()
}
func (v *Try1) Pred(x int) int {
	if x < 0 {
		return v.Max()
	}
	xHi, xLo := v.split(x)

	loPred := (*v.local[xHi]).Pred(int(xLo))
	if loPred != -1 {
		return int(xHi)*v.q + loPred
	}

	hiPred := v.global.Pred(int(xHi))
	if hiPred == -1 {
		return -1
	}

	return hiPred*v.q + (*v.local[hiPred]).Max()
}
func (v *Try1) Delete(x int) {
	xHi, xLo := v.split(x)
	(*v.local[xHi]).Delete(int(xLo))
	if (*v.local[xHi]).Min() == -1 {
		v.global.Delete(int(xHi))
	}
}
func (v *Try1) Min() int {
	hiMin := v.global.Min()
	if hiMin == -1 {
		return -1
	}

	return hiMin*v.q + (*v.local[hiMin]).Min()
}
func (v *Try1) Max() int {
	hiMax := v.global.Max()
	if hiMax == -1 {
		return -1
	}

	return hiMax*v.q + (*v.local[hiMax]).Max()
}
func (v *Try1) Member(x int) bool {
	xHi, xLo := v.split(x)
	hiMem := v.global.Member(int(xHi))
	if !hiMem {
		return false
	}

	return (*v.local[xHi]).Member(int(xLo))
}

func (v *Try1) split(x int) (hi, lo uint64) {
	x2 := uint64(x)

	hi = x2 >> v.q_
	lo = x2 & (^uint64(0) >> (64 - v.q_))
	return
}
