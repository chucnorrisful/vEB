package vEB

// LLPrioQ is a very simple implementation of the PrioQ interface using a linked list.
type LLPrioQ struct {
	start, end *Link
}
type Link struct {
	int
	prev, next *Link
}

func (v *LLPrioQ) Init(u int, fullInit bool) {

}
func (v *LLPrioQ) Insert(x int) {
	xIns := &Link{x, nil, nil}
	if v.start == nil {
		v.start = xIns
		v.end = xIns
		return
	}

	var prev *Link
	curr := v.start
	for curr != nil && curr.int < x {
		prev = curr
		curr = curr.next
	}
	if curr == nil && prev == nil {
		return
	}
	if curr != nil && curr.int == x {
		return
	}

	if prev == nil {
		v.start = xIns
		xIns.next = curr
		curr.prev = xIns
		return
	}
	if curr == nil {
		v.end = xIns
		prev.next = xIns
		xIns.prev = prev
		return
	}

	curr.prev = xIns
	xIns.prev = prev
	xIns.next = curr
	prev.next = xIns
}
func (v *LLPrioQ) Delete(x int) {
	curr := v.start
	for curr != nil && curr.int < x {
		curr = curr.next
	}
	if curr == nil || curr.int > x {
		// not found
		return
	}
	if curr.prev == nil {
		v.start = curr.next
		if v.start == nil {
			return
		}
		curr.next.prev = nil
		return
	}
	if curr.next == nil {
		v.end = curr.prev
		curr.prev.next = nil
		return
	}
	curr.prev.next = curr.next
	curr.next.prev = curr.prev
	return
}
func (v *LLPrioQ) Succ(x int) int {
	if v.start == nil {
		return -1
	}
	curr := v.start
	for curr != nil && curr.int <= x {
		curr = curr.next
	}
	if curr == nil {
		return -1
	}
	return curr.int
}
func (v *LLPrioQ) Pred(x int) int {
	if v.end == nil {
		return -1
	}
	curr := v.end
	for curr != nil && curr.int >= x {
		curr = curr.prev
	}
	if curr == nil {
		return -1
	}
	return curr.int
}
func (v *LLPrioQ) Min() int {
	if v.start == nil {
		return -1
	}
	return v.start.int
}
func (v *LLPrioQ) Max() int {
	if v.end == nil {
		return -1
	}
	return v.end.int
}
func (v *LLPrioQ) Member(x int) bool {

	curr := v.start
	for curr != nil && curr.int <= x {
		if curr.int == x {
			return true
		}
		curr = curr.next
	}

	return false
}
