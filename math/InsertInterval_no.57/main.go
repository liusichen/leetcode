package main

import "fmt"

func main() {
	a := []Interval{
		Interval{1, 2},
		Interval{3, 5},
		Interval{6, 8},
	}
	fmt.Println(insert(a, Interval{2, 5}))
}

//Interval ...
type Interval struct {
	Start int
	End   int
}

func insert(intervals []Interval, newInterval Interval) []Interval {
	start := newInterval.Start
	end := newInterval.End
	var res []Interval
	i := 0
	for i < len(intervals) {
		if start <= intervals[i].End {
			if end < intervals[i].Start {
				break
			}
			if start > intervals[i].Start {
				start = intervals[i].Start
			}
			if end < intervals[i].End {
				end = intervals[i].End
			}
		} else {
			res = append(res, intervals[i])
		}
		i++
	}
	res = append(res, Interval{
		Start: start,
		End:   end,
	})
	res = append(res, intervals[i:]...)
	return res
}
