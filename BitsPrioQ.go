package vEB

import (
	"fmt"
	"math"
	"math/bits"
)

// ArrPrioQ is a very simple implementation of the PrioQ interface using a sorted array.
// It's not very performant, but serves as a baseline to compare against.
type BitsPrioQ struct {
	data []uint64
	mask uint64
	u    int
}

func (v *BitsPrioQ) Init(u int, fullInit bool) {
	a := int(math.Exp2(math.Ceil(math.Log2(float64(u) / 64))))
	v.data = make([]uint64, a)
	v.mask = 63
	v.u = u
}

func (v *BitsPrioQ) Insert(x int) {
	ind := x >> 6
	if ind >= len(v.data) {
		fmt.Println("Cannot insert numbers bigger than", v.u)
		return
	}
	v.data[ind] |= 1 << (uint64(x) & v.mask)
}

func (v *BitsPrioQ) Delete(x int) {
	if !v.Member(x) {
		return
	}
	v.data[x>>6] ^= 1 << (uint64(x) & v.mask)
}

func (v *BitsPrioQ) Succ(x int) int {
	if x < 0 {
		if v.Member(0) {
			return 0
		}
		return v.Succ(0)
	}
	ind := x >> 6
	u := v.data[ind]
	low := uint64(x) & v.mask
	if u >= 1<<(low+1) {
		u = u & (math.MaxUint64 << (low + 1))
		h := x >> 6 << 6
		t := bits.TrailingZeros64(u)
		return h + t
	}
	ind++

	for ind < len(v.data) {
		u = v.data[ind]
		if u != 0 {
			h := ind << 6
			t := bits.TrailingZeros64(u)
			return h + t
		}
		ind++
	}

	return -1
}
func (v *BitsPrioQ) Pred(x int) int {
	if x < 0 {
		return -1
	}
	if x > v.u-1 {
		return v.Max()
	}
	ind := x >> 6
	u := v.data[ind]
	low := uint64(x) & v.mask
	u &= (1 << low) - 1 // mask: only lower bits
	if u > 0 {
		h := x >> 6 << 6
		t := 64 - bits.LeadingZeros64(u) - 1
		return h + t
	}
	ind--

	for ind >= 0 {
		u = v.data[ind]
		if u != 0 {
			h := ind << 6
			t := 64 - bits.LeadingZeros64(u) - 1
			return h + t
		}
		ind--
	}

	return -1
}

func (v *BitsPrioQ) Min() int {
	ind := 0
	var u uint64
	for ind < len(v.data) {
		u = v.data[ind]
		if u != 0 {
			h := ind << 6
			t := bits.TrailingZeros64(u)
			return h + t
		}
		ind++
	}
	return -1
}

func (v *BitsPrioQ) Max() int {
	ind := len(v.data) - 1
	var u uint64
	for ind >= 0 {
		u = v.data[ind]
		if u != 0 {
			h := ind << 6
			t := 64 - bits.LeadingZeros64(u) - 1
			return h + t
		}
		ind--
	}
	return -1
}
func (v *BitsPrioQ) Member(x int) bool {
	return v.data[x>>6]&(1<<(uint64(x)&v.mask)) != 0
}
