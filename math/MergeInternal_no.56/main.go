package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []Interval{
		Interval{1, 4},
		Interval{0, 2},
		Interval{3, 5},
	}
	fmt.Println(merge(a))
}

//Interval ...
type Interval struct {
	Start int
	End   int
}

//Print ...
func (i Interval) Print() string {
	return fmt.Sprintf("start=%d,end=%d\n", i.Start, i.End)
}

//ByStart ...
type ByStart []Interval

func (b ByStart) Len() int           { return len(b) }
func (b ByStart) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByStart) Less(i, j int) bool { return b[i].Start < b[j].Start }

func merge(intervals []Interval) []Interval {
	var res []Interval

	if len(intervals) <= 0 || intervals == nil {
		return nil
	}
	sort.Sort(ByStart(intervals))
	res = append(res, intervals[0])
	resNum := 1
	for i := 1; i < len(intervals); i++ {
		fmt.Println(res)
		tmp := res[resNum-1]
		if tmp.End < intervals[i].Start {
			res = append(res, intervals[i])
			resNum++
		} else if tmp.End >= intervals[i].End {
			continue
		} else {
			res[resNum-1].End = intervals[i].End
		}
	}
	return res

}
