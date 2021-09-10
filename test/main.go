package main

import (
	"fmt"
	"github.com/chucnorrisful/vEB"
)

// main provides some example code for using a vEB tree.
func main() {

	var v vEB.PrioQ = vEB.InitVEB(2 << 26, true)

	v.Insert(1)
	s := v.Succ(0)
	fmt.Printf("successor of 0 is %v.\n", s)

	s = v.Succ(1)
	fmt.Printf("successor of 1 is %v; this means it has no successor.\n", s)

	test := []int{4, 3, 100, 200423, 3492939, 70}
	for _,t := range test {
		v.Insert(t)
	}

	// successor works also for elements, which are not part of the tree
	s = v.Succ(10005)
	fmt.Printf("successor of 10005 is %v.\n", s)

	//this is how delete works
	v.Delete(1)
	v.Delete(3)
	v.Delete(4)
	v.Delete(100)

	//this is how you find the minimum element
	s = v.Succ(-1)
	fmt.Printf("successor of -1 is the minimum stored element: %v.\n", s)

	//this is how you check if a number is a member of the tree:
	maybeMember := 70
	s = v.Succ(maybeMember - 1)
	if s == maybeMember {
		fmt.Printf("%v was found inside the tree.", maybeMember)
	} else {
		fmt.Printf("%v was not found inside the tree.", maybeMember)
	}
}
