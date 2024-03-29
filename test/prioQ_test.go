package main

import (
	"fmt"
	"github.com/chucnorrisful/vEB"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// todo: Add tests for overflows and underflows (succ, pred, insert, delete, member)

var algosTst = []algo{
	{"ll", func() vEB.PrioQ { return &vEB.LLPrioQ{} }},
	{"arr", func() vEB.PrioQ { return &vEB.ArrPrioQ{} }},
	{"bits", func() vEB.PrioQ { return &vEB.BitsPrioQ{} }},
	{"try0", func() vEB.PrioQ { return &vEB.Try0{} }},
	{"try1", func() vEB.PrioQ { return &vEB.Try1{} }},
	{"try2", func() vEB.PrioQ { return &vEB.Try2{} }},
	{"try3", func() vEB.PrioQ { return &vEB.Try3{} }},
	//{"try4", func() vEB.PrioQ { return &vEB.Try4{} }},
	{"v0", func() vEB.PrioQ { return &vEB.V0{} }},
	{"v1", func() vEB.PrioQ { return &vEB.V1{} }},
}

func TestPrioQ(t *testing.T) {
	for _, alg := range algosTst {
		name := alg.name
		v := alg.gen()
		t.Run(fmt.Sprintf("%v_interface", name), func(t *testing.T) {
			v.Init(1000, false)
			var s = -1

			v.Insert(1)

			s = v.Succ(0)
			if s != 1 {
				t.Errorf("%s: succ should have been 1 but was %v", name, s)
			}
			s = v.Succ(1)
			if s != -1 {
				t.Errorf("%s: succ should have been -1 but was %v", name, s)
			}
			s = v.Pred(0)
			if s != -1 {
				t.Errorf("%s: pred should have been -1 but was %v", name, s)
			}
			s = v.Pred(3)
			if s != 1 {
				t.Errorf("%s: pred should have been 1 but was %v", name, s)
			}

			v.Insert(4)
			v.Insert(3)
			v.Insert(100)

			if !v.Member(4) {
				t.Errorf("%s: 4 should have been a member", name)
			}
			if v.Member(5) {
				t.Errorf("%s: 5 shouldn't have been a member", name)
			}

			if v.Max() != 100 {
				t.Errorf("%s: max should have been %d but was %d", name, 100, v.Max())
			}
			if v.Min() != 1 {
				t.Errorf("%s: min should have been %d but was %d", name, 1, v.Min())
			}

			v.Delete(1)
			if v.Member(1) {
				t.Errorf("%s: 1 shouldn't have been a member", name)
			}

			s = v.Succ(0)
			if s != 3 {
				t.Errorf("%s: succ should have been 3 but was %v", name, s)
			}
			s = v.Succ(4)
			if s != 100 {
				t.Errorf("%s: succ should have been 100 but was %v", name, s)
			}
			s = v.Pred(3)
			if s != -1 {
				t.Errorf("%s: pred should have been -1 but was %v", name, s)
			}
			s = v.Pred(4)
			if s != 3 {
				t.Errorf("%s: pred should have been 3 but was %v", name, s)
			}

			v.Delete(3)
			v.Delete(4)
			v.Delete(100)

			s = v.Succ(-1)
			if s != -1 {
				t.Errorf("%s: succ (-1) should have been -1 but was %v", name, s)
			}
			s = v.Pred(-1)
			if s != -1 {
				t.Errorf("%s: succ (-1) should have been -1 but was %v", name, s)
			}
		})
	}

	u := 32
	rng := rand.Perm(u)
	ins := rng[:int(float64(len(rng))*0.7)]
	del := ins[len(ins)/4 : len(ins)*3/4]
	for _, alg := range algosTst {
		name := alg.name
		v := alg.gen()
		t.Run(fmt.Sprintf("%v_%d_load", name, u), func(t *testing.T) {
			PrioQInitTask(v, u, true)
			PrioQLoadTask(v, rng, ins, del)
		})
	}

	u = 10_000
	rng = rand.Perm(u)
	ins = rng[:int(float64(len(rng))*0.7)]
	del = ins[len(ins)/4 : len(ins)*3/4]
	for _, alg := range algosTst {
		name := alg.name
		v := alg.gen()
		t.Run(fmt.Sprintf("%v_%d_load", name, u), func(t *testing.T) {
			PrioQInitTask(v, u, true)
			PrioQLoadTask(v, rng, ins, del)
		})
	}

	// Do LoadTest and compare results
	rems := make(map[string][]int)
	for _, alg := range algosTst {
		name := alg.name
		v := alg.gen()
		rem := make([]int, 0, len(ins))
		var x int
		for {
			x = v.Succ(-1)
			if x == -1 {
				break
			}
			rem = append(rem, x)
			v.Delete(x)
		}
		rems[name] = rem
	}
	l, lN := -1, ""
	for name, ints := range rems {
		if l != -1 {
			if l != len(ints) {
				t.Errorf("%s output len does not align with %s: %d != %d", name, lN, len(ints), l)
			}
			if !reflect.DeepEqual(rems[name], rems[lN]) {
				t.Errorf("%s output is not equal to %s", name, lN)
			}
		}
		l = len(ints)
		lN = name
	}
}

type algo struct {
	name string
	gen  func() vEB.PrioQ
}

var algos = []algo{
	//{"ll", func() vEB.PrioQ { return &vEB.LLPrioQ{} }},
	//{"arr", func() vEB.PrioQ { return &vEB.ArrPrioQ{} }},
	{"bits", func() vEB.PrioQ { return &vEB.BitsPrioQ{} }},
	//{"try0", func() vEB.PrioQ { return &vEB.Try0{} }},
	//{"try1", func() vEB.PrioQ { return &vEB.Try1{} }},
	//{"try2", func() vEB.PrioQ { return &vEB.Try2{} }},
	//{"try3", func() vEB.PrioQ { return &vEB.Try3{} }},
	{"try4", func() vEB.PrioQ { return &vEB.Try4{} }},
	//{"v0", func() vEB.PrioQ { return &vEB.V0{} }},
	//{"v1", func() vEB.PrioQ { return &vEB.V1{} }},
	//{"std", func() vEB.PrioQ { return nil }},
}
var sizes = []int{
	//100,
	//1000,
	//10_000,
	//100_000,
	//1_000_000,
	1_000_000,
	//100_000_000,
}

func BenchmarkSortAll(b *testing.B) {
	for _, v := range sizes {
		rngMain := rand.Perm(v)
		for _, algo := range algos {
			a := algo.gen()
			PrioQInitTask(a, v, true)

			b.Run(fmt.Sprintf("size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					b.StopTimer()
					rng := rngMain
					b.StartTimer()
					PrioQSortTask(a, rng)
				}
			})
		}
	}
}
func BenchmarkLoadAll(b *testing.B) {
	for _, v := range sizes {
		rngMain := rand.Perm(v)
		insMain := rngMain[:int(float64(len(rngMain))*0.7)]
		delMain := insMain[len(insMain)/4 : len(insMain)*3/4]

		for _, algo := range algos {
			a := algo.gen()
			if a == nil {
				continue
			}
			PrioQInitTask(a, v, true)

			b.Run(fmt.Sprintf("size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					rng, ins, del := rngMain, insMain, delMain

					PrioQLoadTask(a, rng, ins, del)
				}
			})
		}
	}
}

func BenchmarkInitAll(b *testing.B) {
	for _, v := range sizes {
		for _, algo := range algos {
			b.Run(fmt.Sprintf("init_size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					a := algo.gen()
					a.Init(v, true)
				}
			})
		}
	}
}
func BenchmarkMemberAll(b *testing.B) {
	for _, v := range sizes {
		rng := rand.Perm(v)
		ins := rng[:int(float64(len(rng))*0.7)]

		for _, algo := range algos {
			a := algo.gen()
			if a == nil {
				continue
			}
			PrioQInitTask(a, v, true)
			for x := range ins {
				a.Insert(x)
			}

			b.Run(fmt.Sprintf("size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					PrioQMemberTask(a, rng)
				}
			})
		}
	}
}
func BenchmarkSuccAll(b *testing.B) {
	for _, v := range sizes {
		rng := rand.Perm(v)
		ins := rng[:int(float64(len(rng))*0.7)]

		for _, algo := range algos {
			a := algo.gen()
			if a == nil {
				continue
			}
			PrioQInitTask(a, v, true)
			for x := range ins {
				a.Insert(x)
			}

			b.Run(fmt.Sprintf("size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					PrioQSuccTask(a, rng)
				}
			})
		}
	}
}
func BenchmarkInsertAll(b *testing.B) {
	for _, v := range sizes {
		rng := rand.Perm(v)

		for _, algo := range algos {
			a := algo.gen()
			if a == nil {
				continue
			}
			PrioQInitTask(a, v, true)

			b.Run(fmt.Sprintf("size_%d_algo_%v", v, algo.name), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					PrioQInsertTask(a, rng)
				}
			})
		}
	}
}

func PrioQInitTask(pq vEB.PrioQ, u int, fullInit bool) {
	if pq == nil {
		return
	}

	pq.Init(u, fullInit)
}
func PrioQLoadTask(pq vEB.PrioQ, rng, ins, del []int) {
	if pq == nil {
		return
	}

	for x := range ins {
		pq.Insert(x)
	}

	for x := range rng {
		pq.Succ(x)
	}

	for x := range rng {
		pq.Member(x)
	}

	for x := range del {
		pq.Delete(x)
	}

	for x := range rng {
		pq.Succ(x)
	}
}
func PrioQMemberTask(pq vEB.PrioQ, rng []int) {
	for x := range rng {
		pq.Member(x)
	}
}
func PrioQInsertTask(pq vEB.PrioQ, rng []int) {
	for x := range rng {
		pq.Insert(x)
	}
}
func PrioQSuccTask(pq vEB.PrioQ, rng []int) {
	for x := range rng {
		pq.Succ(x)
	}
}
func PrioQSortTask(pq vEB.PrioQ, rng []int) []int {
	if pq == nil {
		// std.sort default case to compare
		sort.Ints(rng)
		return rng
	}

	for x := range rng {
		pq.Insert(x)
	}
	res := make([]int, len(rng))

	at := -1
	for i := 0; i < len(rng); i++ {
		at = pq.Succ(at)
		res[i] = at
	}

	return res
}
