package main 

import(
	"fmt"
	"sort"
)

//56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，
// 如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
func MergeIntervals(i [][]int) [][]int {
	if len(i) <= 1 {
		return i
	}

	sort.Slice(i, func(a, b int) bool {
		return i[a][0] < i[b][0]
	})

	i2 := make([][]int, 0)
	for j:=0;j<len(i);j++{
		if len(i2) == 0 {
			i2 = append(i2, i[j])
			continue
		}

		last := i2[len(i2)-1]
		if last[1] >= i[j][0] {
			last[1] = max(last[1], i[j][1])
		} else {
			i2 = append(i2, i[j])
		}
	}

	return i2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(MergeIntervals([][]int{{8,10},{1,3},{15,18},{2,6}, {5, 10}}))
}