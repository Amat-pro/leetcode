package greed

// findContentChildren - 0455 - 分发饼干
func findContentChildren(g []int, children []int) int {

	// 用大饼干喂胃口大的孩子

	result := 0

	gIndex := len(g) - 1
	for gIndex >= 0 {
		for child := len(children)-1; child >= 0; child-- {

			if g[gIndex] >= children[child] { // 满足胃口，执行投喂
				result++
				gIndex--
			}

			// 否则，这块饼干寻找下一个胃口大的小孩

		}
	}

	return result

}