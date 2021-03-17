package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	vIn    int
	group  int
	item   int
	next   []int
}

/*
* 题目来源：https://leetcode-cn.com/problems/sort-items-by-groups-respecting-dependencies/
*/
func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	m = initGroup(m, group)

	// 每个组做的项目list
	gItMap := make(map[int][]int, 0)
	for k, _:= range group {
		g := group[k]
		if _, ok := gItMap[g]; ! ok {
			gItMap[g] = make([]int, 0)
		}
		gItMap[g] = append(gItMap[g], k)
	}
	// fmt.Println(gItMap)

	// 构建项目依赖图
	itemGraph := buildItemGraph(n, group, beforeItems)
	//for k, _:= range itemGraph {
	//	printNode(itemGraph[k])
	//}
	// fmt.Println()

	// 判断是否有环
	//bCircle := topologyCheckCircle(n, itemGraph)
	//if bCircle {
	//	fmt.Println("check circle return true")
	//	return []int{}
	//}

	// 构建group依赖图
	groupGraph := buildGroupGraph(n, m, group, beforeItems)
	//for k, _:= range groupGraph {
	//	printNode(groupGraph[k])
	//}

	// 先遍历group，根据group所做的项目，去遍历项目
	result := make([]int, 0)
	gt := list.New()
	for k:=0; k < m; k++ {
		if groupGraph[k] != nil && groupGraph[k].vIn == 0 {
			gt.PushBack(k)
		}
	}
	for ;gt.Len() > 0; {
		elem := gt.Back()
		gt.Remove(elem)

		k := elem.Value.(int)
		//fmt.Println("group: ", k)
		for _, next := range groupGraph[k].next {
			groupGraph[next].vIn--
			if groupGraph[next].vIn == 0 {
				gt.PushBack(next)
			}
		}

		// 遍历group所做的项目
		res := itemTravel(k, group, gItMap[k], itemGraph)
		result = append(result, res...)
	}

	if len(result) != n {
		return []int{}
	}

	return result
}

func itemTravel(g int, group []int, its[]int, itGraph []*Node) []int {
	// fmt.Print("items: ", its)
	lt := list.New()
	for _, k := range its {
		if itGraph[k].vIn == 0 {
			lt.PushBack(k)
		}
	}

	res := make([]int, 0)
	for ;lt.Len() > 0; {
		elem := lt.Back()
		lt.Remove(elem)

		k := elem.Value.(int)
		// fmt.Println(" item ", k)
		for _, next := range itGraph[k].next {
			itGraph[next].vIn--
			if itGraph[next].vIn == 0 && group[next] == g {
				lt.PushBack(next)
			}
		}

		res = append(res, k)
	}

	return res
}

// 构建group依赖图
func buildGroupGraph(n int, m int, group []int, beforeItems [][]int) []*Node {
	gNodes := make([]*Node, m, m)
	for k, _:= range group {
		if gNodes[group[k]] != nil {
			continue
		}

		node := makeNode(group[k], k)
		gNodes[group[k]] = node
	}

	for k, deps := range beforeItems {
		if len(deps) == 0 {
			continue
		}

		for _, dep := range deps {
			g := group[dep]
			if group[k] != g {
				gNodes[g].next = append(gNodes[g].next, group[k])
				gNodes[group[k]].vIn++
			}
		}
	}

	return gNodes
}

// 项目依赖图
func buildItemGraph(n int, group []int, beforeItems [][]int) []*Node {
	itemNodes := make([]*Node, n, n)
	for k, _:= range group {
		node := makeNode(group[k], k)
		itemNodes[k] = node
	}

	for k, deps := range beforeItems {
		if len(deps) == 0 {
			continue
		}

		// next->k, k的入度加1
		for _, dep := range deps {
			itemNodes[dep].next = append(itemNodes[dep].next, k)
			itemNodes[k].vIn++
		}
	}

	return itemNodes
}

func makeNode(group int, item int) *Node {
	return &Node{
		vIn:   0,
		group: group,
		item:  item,
		next:  make([]int, 0),
	}
}

func printNode(n *Node) {
	fmt.Printf("group=%d, item=%d, vIn=%d，next=%v \n", n.group, n.item, n.vIn, n.next)
}

// 初始化没有分配任务的项目
func initGroup(m int, group []int) int {
	for k, _:= range group {
		if group[k] == -1 {
			group[k] = m
			m++
		}
	}

	return m
}


func test1() {
	n := 8
	m := 2
	group := []int{-1,-1,1,0,0,1,0,-1}
	beforeItems := [][]int{{},{6},{5},{6},{3,6},{},{},{}}
	result := sortItems(n, m, group, beforeItems)
	fmt.Println(result)
}

func test2() {
	n := 8
	m := 2
	group := []int{-1,-1,1,0,0,1,0,-1}
	beforeItems := [][]int{{},{6},{5},{6},{3,6},{},{3},{}}
	result := sortItems(n, m, group, beforeItems)
	fmt.Println(result)
}

func test3() {
	n := 5
	m := 5
	group := []int{2,0,-1,3,0}
	beforeItems := [][]int{{2,1,3},{2,4},{},{},{}}
	result := sortItems(n, m, group, beforeItems)
	fmt.Println(result)
}

func main() {
	test3()
}

