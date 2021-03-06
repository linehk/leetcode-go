package duplicationInArrayNoEdit

// 计数法（不可得全部）
// O(n) O(n)
func duplicationInArrayNoEdit1(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	s := make([]int, len(nums))
	for _, v := range nums {
		if s[v] == v {
			return v, true
		}
		s[v] = v
	}

	return 0, false
}

// 计数法（两次扫描，可得全部）
// O(2n) O(n)
func duplicationInArrayNoEdit2(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	s := make([]int, len(nums))
	for _, v := range nums {
		s[v]++
	}
	for i, v := range s {
		if v > 1 {
			return i, true
		}
	}

	return 0, false
}

// 二分法
// O(nlgn) O(1)
func duplicationInArrayNoEdit3(nums []int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}

	start := 1
	end := len(nums) - 1
	for start <= end {
		mid := start + (end-start)>>1
		count := countRange(nums, start, mid)
		if start == end {
			if count > 1 {
				return start, true
			} else {
				// return 0, false
				break
			}
		}

		if count > (mid - start + 1) {
			end = mid
		} else {
			start = mid + 1
		}
	}

	return 0, false
}

// 统计数组元素在 [start, end] 区间出现过的次数
func countRange(nums []int, start, end int) int {
	count := 0
	for _, v := range nums {
		if start <= v && v <= end {
			count++
		}
	}
	return count
}
