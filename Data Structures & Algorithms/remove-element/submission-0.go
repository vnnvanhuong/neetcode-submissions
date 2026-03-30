// problem
// need to set k elements not equal to val

// logical solution
// k=0
// loop through nums, if nums[i] != val
// set nums[k]=num[i], increase k
// return k

// dry-run
// 1, 1, 2, 3, 4
// 0: 1, 1, 2, 3, 4
// 1: 1, 1, 2, 3, 4
// 2: 2, 1, 2, 3, 4
// 2: 2, 3, 2, 3, 4
// 2: 2, 3, 4, 3, 4 -> k=3

func removeElement(nums []int, val int) int {
    k := 0
    for i := range nums {
        if nums[i] != val {
            nums[k] = nums[i]
            k++
        }
    }

    return k
}

