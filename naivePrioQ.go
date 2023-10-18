package vEB

// NaivePrioQ is a very simple implementation of the PrioQ interface using a sorted array.
// It's not very performant, but serves as a baseline to compare against.
type NaivePrioQ struct {
	data []int
}

func (v *NaivePrioQ) Init(u int, fullInit bool) {
	v.data = make([]int, 0)
}

func (v *NaivePrioQ) Insert(x int) {
	var small = -1
	for i, d := range v.data {
		if d < x {
			small = i
		} else {
			break
		}
	}

	v.data = append(v.data, 0)
	if len(v.data)-small > 2 {
		copy(v.data[small+2:], v.data[small+1:])
	}
	v.data[small+1] = x
}
func (v *NaivePrioQ) Delete(x int) {
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
func (v *NaivePrioQ) Succ(x int) int {
	for _, d := range v.data {
		if d > x {
			return d
		}
	}
	return -1
}
func (v *NaivePrioQ) Pred(x int) int {
	for i := len(v.data) - 1; i >= 0; i-- {
		if v.data[i] < x {
			return v.data[i]
		}
	}
	return -1
}
func (v *NaivePrioQ) Min() int {
	if len(v.data) == 0 {
		return -1
	}

	return v.data[0]
}
func (v *NaivePrioQ) Max() int {
	if len(v.data) == 0 {
		return -1
	}

	return v.data[len(v.data)-1]
}
func (v *NaivePrioQ) Member(x int) bool {
	for _, d := range v.data {
		if x == d {
			return true
		}
	}
	return false
}
