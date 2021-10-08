package ufsm

func Merge(x int, y int, sets map[int]int, rank map[int]int) {
	x = Find(x, sets)
	y = Find(y, sets)
	if x == y {
		return
	}

	if rank[x] > rank[y] {
		sets[y] = x
	} else {
		sets[x] = y
		rank[y]++
	}
}
