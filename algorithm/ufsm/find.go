package ufsm

func Find(x int, sets map[int]int) int {
	// 首先找到x的祖先
	p := x
	for ; sets[p] != p;  {
		p = sets[p]
	}

	// 然后把x的父亲、爷爷等等，都设置成x的祖先
	var t int
	for ; x != p ; {
		t = sets[x] // 保存x的父亲
		sets[x] = p // 把父亲也设置成最早的祖先
		x = t // 继续设置 父亲
	}

	return x
}