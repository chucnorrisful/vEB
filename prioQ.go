package vEB

//todo: extend to also feature Pred, Min, Max and Member

type PrioQ interface {
	Insert(x int)
	Delete(x int)
	Succ(x int) int
	Pred(x int) int
	Member(x int) bool
	Min() int
	Max() int

	Init(u int, fullInit bool)
}
