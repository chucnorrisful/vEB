package vEB

//todo: extend to also feature Pred, Min, Max and Member

type PrioQ interface {
	Init(u int, fullInit bool)
	Insert(x int)
	Delete(x int)
	Succ(x int) int
}
