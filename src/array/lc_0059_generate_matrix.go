package array

// generateMatrix 螺旋矩阵
func generateMatrix(n int) [][]int {

	if n <= 0 {
		return nil
	}

	nums := make([][]int, 0, n)
	for i := 0; i < n ; i ++ {
		nums = append(nums, make([]int, n))
	}

	num := 1 // 数字计数器
	count := 1  // 记录圈数

	startx := 0
	starty := 0

	// x, y 表示变动的坐标位置 
	var x int
	var y int


	for count <= n / 2 {
		// top边  [, )
		// x不变, y递增
		for y = starty; y < n - count; y++ {
			nums[startx][y] = num
			num++
		}

		// right边 [, )
		// y不变, x递增
		for x = startx; x < n - count; x++ {
			nums[x][y] = num
			num++
		}
		
		// bottom边 [, )
		// x不变, y递减
		for ; y > starty ; y-- {
			nums[x][y] = num
			num++
		}

		// left边 [, )
		// y不变, x递减
		for ; x > startx ; x-- {
			nums[x][y] = num
			num++
		}

	
		count++
		startx++
		starty++

	}

	if n % 2 == 1 {  // 奇数
		idx := n /2 
		nums[idx][idx] = n * n
	}

	return nums
}