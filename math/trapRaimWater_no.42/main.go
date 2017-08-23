package main

import "fmt"

func main() {
	a := []int{0, 5, 1, 3, 4, 2, 0, 6, 0, 4, 0, 2, 0}
	fmt.Println(trap(a))
}

type trapPos struct {
	Height int
	Pos    int
}

func trap(height []int) int {
	if height == nil || len(height) <= 1 {
		return 0
	}

	dp := make([]int, len(height))
	stack := make([]trapPos, len(height))
	top := 0
	dp[0] = 0
	heightPos := 0
	for height[heightPos] == 0 {
		dp[heightPos] = 0
		heightPos++
	}
	if heightPos == len(height)-1 {
		return 0
	}
	stack[0] = trapPos{
		Height: height[heightPos],
		Pos:    heightPos,
	}
	top++
	dp[heightPos] = 0

	for i := heightPos + 1; i < len(height); i++ {

		if height[i] == 0 {
			dp[i] = dp[i-1]
			continue
		}
		if i == stack[top-1].Pos+1 && height[i] < stack[top-1].Height {
			dp[i] = dp[i-1]
			stack[top] = trapPos{
				Height: height[i],
				Pos:    i,
			}
			top++
			continue
		}

		tmpHeight := height[i] - height[i-1]
		preHeight := 0
		tmpDP := 0
		for top != 0 {
			fmt.Printf("height[i]=%d,stack[%d].Height=%d\n", height[i], top-1, stack[top-1].Height)
			if height[i] < stack[top-1].Height {
				tmpHeight = height[i] - preHeight
				tmpDP += tmpHeight * (i - stack[top-1].Pos - 1)
				stack[top] = trapPos{
					Height: height[i],
					Pos:    i,
				}
				top++
				break
			}

			if top == 1 {
				if stack[top-1].Pos != i-1 {
					tmpHeight = stack[top-1].Height - preHeight
					tmpDP += tmpHeight * (i - stack[top-1].Pos - 1)
				}
				top--
				break
			}
			tmpTrap := stack[top-1]
			top--
			if tmpTrap.Pos == i-1 {
				if height[i] > stack[top-1].Height {
					tmpHeight = stack[top-1].Height - tmpTrap.Height
				}
			} else {
				tmpHeight = tmpTrap.Height - preHeight
			}
			fmt.Printf("tmpHeight=%d,top=%d\n", tmpHeight, top)
			tmpDP += tmpHeight * (i - tmpTrap.Pos - 1)
			preHeight = tmpTrap.Height
		}
		fmt.Printf("i=%d, tmpDP=%d, dp[i-1]=%d\n", i, tmpDP, dp[i-1])
		dp[i] = dp[i-1] + tmpDP
		if top == 0 {
			stack[top] = trapPos{
				Height: height[i],
				Pos:    i,
			}
			top++
		}
	}
	return dp[len(height)-1]
}
