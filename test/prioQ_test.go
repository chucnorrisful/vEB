package main

import (
	"fmt"
	"github.com/chucnorrisful/vEB"
	"math/rand"
	"testing"
)

func TestPrioQ(t *testing.T) {
	var structsToTest = map[string]vEB.PrioQ{
		"naive": vEB.InitNaivePrioQ(),
		"vEB":   vEB.InitVEB(1024),
	}

	for name := range structsToTest {
		var v vEB.PrioQ = structsToTest[name]
		var s = -1

		v.Insert(1)

		s = v.Succ(0)
		if s != 1 {
			t.Errorf("succ should have been 1 but was %v", s)
		}
		s = v.Succ(1)
		if s != -1 {
			t.Errorf("succ should have been -1 but was %v", s)
		}

		v.Insert(4)
		v.Insert(3)
		v.Insert(100)

		v.Delete(1)

		s = v.Succ(0)
		if s != 3 {
			t.Errorf("succ should have been 3 but was %v", s)
		}
		s = v.Succ(4)
		if s != 100 {
			t.Errorf("succ should have been 100 but was %v", s)
		}

		v.Delete(3)
		v.Delete(4)
		v.Delete(100)

		s = v.Succ(-1)
		if s != -1 {
			t.Errorf("succ (2) should have been -1 but was %v", s)
		}
	}
}

func BenchmarkNaivePrioQ(b *testing.B) {
	// create 100k random numbers
	rngCnt := 10_000
	fmt.Printf("Creating %d random numbers ... ", rngCnt)
	rng := make(map[int]bool, rngCnt)
	var tmp int
	for i := 0; i <rngCnt; i++ {
		tmp = int(rand.Uint32())
		if _,ok := rng[tmp]; !ok {
			rng[tmp] = true
		}
	}
	fmt.Print("done!\n")

	b.StartTimer()
	for i := 0; i < b.N; i++ {

		v := vEB.InitNaivePrioQ()

		for x := range rng {
			v.Insert(x)
		}

		for x := range rng {
			v.Succ(x)
		}

		for x := range rng {
			v.Delete(x)
		}
	}
}
func BenchmarkVEBPrioQ(b *testing.B) {
	// create 100k random numbers
	rngCnt := 10_000
	fmt.Printf("Creating %d random numbers ... ", rngCnt)
	rng := make(map[int]bool, rngCnt)
	var tmp, max int
	for i := 0; i <rngCnt; i++ {
		tmp = int(rand.Uint32())
		if _,ok := rng[tmp]; !ok {
			rng[tmp] = true
			if max < tmp {
				max = tmp
			}
		}
	}
	fmt.Print("done!\n")

	b.StartTimer()
	for i := 0; i < b.N; i++ {

		v := vEB.InitVEB(max)

		for x := range rng {
			v.Insert(x)
		}

		for x := range rng {
			v.Succ(x)
		}

		for x := range rng {
			v.Delete(x)
		}
	}
}
