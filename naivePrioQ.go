package vEB

// NaivePrioQ is a very simple implementation of the PrioQ interface using a sorted array.
// Its not very performant, but serves as a simple benchmark to compare the vEB tree against.
type NaivePrioQ struct {
	data []int
}

func InitNaivePrioQ() PrioQ {
	var v PrioQ = &NaivePrioQ{
		data: make([]int, 0),
	}
	return v
}

func (v	*NaivePrioQ) Insert(x int) {
	var small = -1
	for i, d := range v.data {
		if d < x {
			small = i
		} else {
			break
		}
	}

	v.data = append(v.data, 0)
	if len(v.data) - small > 2 {
		copy(v.data[small+2:], v.data[small+1:])
	}
	v.data[small+1] = x
}
func (v	*NaivePrioQ) Delete(x int) {
	var xInd = -1
	for i, d := range v.data {
		if d == x {
			xInd = i
			break
		}
	}
	if xInd > -1 {
		v.data = append(v.data[:xInd], v.data[xInd+1:]...)
	}
}
func (v	*NaivePrioQ) Succ(x int) int {
	for _, d := range v.data {
		if d > x {
			return d
		}
	}
	return -1
}

