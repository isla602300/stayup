package myFunc

import (
	"fmt"
	"math"
	"sort"
)

type Tree struct {
	X, Y int
}

func distance(a, b Tree) float64 {
	return math.Sqrt(float64((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

// myFunc returns the length of rope
func myFunc(forest []Tree) (length int) {
	// implement here
	outTrees := outerTrees(forest)
	index := 0
	edge := len(outTrees)
	checkedPointsMap := make(map[int]int, len(outTrees))
	for index < edge {
		disMin := -1.0
		for i := 0; i < len(outTrees); i++ {
			if i != index && disMin == -1.0 {
				disMin = distance(outTrees[i], outTrees[index])
				continue
			}
			if i != index {
				if distance(outTrees[i], outTrees[index]) < disMin {
					disMin = distance(outTrees[i], outTrees[index])
				}
			}
		}
		length = length + int(disMin)
		checkedPointsMap[index] = 1
		fmt.Println(disMin)
		count := 1
		tmpIndex := index
		for count < len(outTrees) {
			index = (tmpIndex + count) % len(outTrees)
			if _, ok := checkedPointsMap[index]; !ok {
				break
			}
			count++
		}
		if count == len(outTrees) {
			return
		}
	}
	return
}

func ConvexHull(p1, p2, p3 Tree) int {
	//S=(x1y2+x2y3+x3y1-x1y3-x2y1-x3y2)/2
	return (p1.X*p2.Y + p2.X*p3.Y + p3.X*p1.Y - p3.X*p2.Y - p2.X*p1.Y - p1.X*p3.Y)
}

func divide(trees []Tree, left, right int, MaxMap map[int]bool) {
	max := 0
	maxPoint := -1
	onLine := make(map[int]bool)
	if left < right {
		for i := left + 1; i < right; i++ {
			v := ConvexHull(trees[left], trees[i], trees[right])
			if v > max {
				max = v
				maxPoint = i
			}
			if v == 0 {
				onLine[i] = true
			}
		}
	} else {
		for i := left - 1; i > right; i-- {
			v := ConvexHull(trees[left], trees[i], trees[right])
			if v > max {
				max = v
				maxPoint = i
			}
			if v == 0 {
				onLine[i] = true
			}
		}
	}

	if maxPoint != -1 {
		MaxMap[maxPoint] = true
		divide(trees, left, maxPoint, MaxMap)
		divide(trees, maxPoint, right, MaxMap)
	} else {
		for k := range onLine {
			MaxMap[k] = true
		}
	}
}

func outerTrees(trees []Tree) []Tree {
	//将树按照横坐标大小排序
	sort.Slice(trees, func(i, j int) bool {
		if trees[i].X == trees[j].X {
			return trees[i].Y < trees[j].Y
		}
		return trees[i].X < trees[j].X
	})

	MaxMap := make(map[int]bool)
	left, right := 0, len(trees)-1
	MaxMap[left] = true
	MaxMap[right] = true
	//不断递归找到两点间的一点使得三点构成的闭包面积最大
	divide(trees, left, right, MaxMap)
	divide(trees, right, left, MaxMap)
	r := make([]Tree, 0, len(trees))
	for k := range MaxMap {
		r = append(r, trees[k])
	}
	return r
}
