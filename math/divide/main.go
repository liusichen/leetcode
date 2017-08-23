package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(15, 4))
}

func divide(dividend int, divisor int) int {
	var dds, dss, sign bool
	if divisor == 0 {
		return math.MaxInt32
	}
	if dividend < 0 {
		dividend = -dividend
		dds = false
	}
	if divisor < 0 {
		divisor = -divisor
		dss = false
	}
	if (!dds && dss) || (dds && !dss) {
		sign = false
	} else {
		sign = true
	}

	if dividend < divisor {
		return 0
	}

	mi := 0
	out := 0
	for (divisor << uint(mi)) <= dividend {
		mi++
	}
	mi--

	for i := mi; i >= 0; i-- {
		if (divisor << uint(i)) <= dividend {
			dividend -= (divisor << uint(i))
			out |= 1 << uint(i)
		}
	}

	if !sign {
		out = -out
	}

	return out

}
