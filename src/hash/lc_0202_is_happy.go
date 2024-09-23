package hash

// isHappy 快乐数
func isHappy(n int) bool {

	set := map[int]bool{}

	sum := n
	for sum != 1 && !set[sum] {
		set[sum] = true
		sum = compute(sum)
	} 

	return sum == 1

}

func compute(n int) int {

	sum := 0
	for n != 0 {
		mod := n % 10
		sum += mod * mod
		n = n / 10
	}

	return sum

}
