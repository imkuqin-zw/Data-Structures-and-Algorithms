package heap

func findKthLargest(nums []int, k int) int {
	tops := make([]int, 0, k)
	for i := 0; i < k; i++ {
		tops = append(tops, nums[i])
		findKthLargest_up(tops, i)
	}
	for i := k; i < len(nums); i++ {
		findKthLargest_push(tops, nums[i])
	}
	return tops[0]
}

func findKthLargest_down(nums []int, i int) {
	j := i
	for {
		j = i<<1 + 1
		if j >= len(nums) {
			break
		}
		if nums[j] > nums[i] {
			j++
			if j >= len(nums) || nums[j] > nums[i] {
				break
			}
		} else {
			if j+1 < len(nums) && nums[j] > nums[j+1] {
				j++
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}
func findKthLargest_up(nums []int, i int) {
	j := i
	for {
		j = (i - 1) / 2
		if j == i || nums[i] > nums[j] {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i = j
	}
}

func findKthLargest_push(nums []int, num int) {
	if num <= nums[0] {
		return
	}
	nums[0] = num
	findKthLargest_down(nums, 0)
}
