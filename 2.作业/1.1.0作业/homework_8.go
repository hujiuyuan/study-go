package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

示例 2：

输入：intervals = [[1,4],[4,5]]
输出：[[1,5]]
解释：区间 [1,4] 和 [4,5] 可被视为重叠区间。
*/

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals)) // 输出：[[1 6] [8 10] [15 18]]

	intervals = [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals)) // 输出：[[1 5]]
}

// merge 合并所有重叠的区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	// 按照区间的起始位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 初始化结果切片
	var ans [][]int
	i := 0
	j := i + 1
	n := len(intervals)

	for j < n {
		if intervals[j][0] <= intervals[i][1] {
			// 如果当前区间与前一个区间重叠，更新前一个区间的结束位置
			intervals[i][1] = max(intervals[j][1], intervals[i][1])
			j++
		} else {
			// 如果不重叠，将前一个区间加入结果切片
			ans = append(ans, intervals[i])
			i = j
		}
	}

	// 添加最后一个区间
	ans = append(ans, intervals[i])

	return ans
}

// max 返回两个整数中的最大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
