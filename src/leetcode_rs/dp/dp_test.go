package dp

import (
	"fmt"
	"testing"
)

func Test_fib(t *testing.T) {
	n := 4
	fmt.Println("fib result is: ", fib(n)) // 3
}

func Test_climbStairs(t *testing.T) {
	n := 3
	fmt.Println("climbStairs result is: ", climbStairs(n)) // 3
}

func Test_minCostClimbingStairs(t *testing.T) {
	cost := []int{10, 15, 20}
	fmt.Println("minCostClimbingStairs result is: ", minCostClimbingStairs(cost)) // 15
	cost2 := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	fmt.Println("minCostClimbingStairs result is: ", minCostClimbingStairs(cost2)) // 6
}

func Test_uniquePaths(t *testing.T) {
	m := 3
	n := 7
	fmt.Println("uniquePaths result is: ", uniquePaths(m, n))
}

func Test_uniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	fmt.Println("uniquePathsWithObstacles result is: ", uniquePathsWithObstacles(obstacleGrid)) // 2
}

func Test_integerBreak(t *testing.T) {
	n := 10
	fmt.Println("整数拆分 - integerBreak result is: ", integerBreak(n)) // 36
}

func Test_bag01(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}

	k := 4

	fmt.Println("bag01 result is: ", bag01(weights, values, k))
	fmt.Println("滚动数组-bag01II result is: ", bag01_II(weights, values, k))
}

func Test_bag(t *testing.T) {
	weights := []int{1, 3, 4}
	values := []int{15, 20, 30}

	k := 4

	fmt.Println("二维数组-bagII result is: ", bag(weights, values, k))
	fmt.Println("滚动数组-bag_II result is: ", bag_II(weights, values, k))
}

