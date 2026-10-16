func longestConsecutive(nums []int) int {
	elements := make(map[int]struct{})

	for _, num := range nums {
		elements[num] = struct{}{}
	}

	var count int

	for k, _ := range elements {
		if _, exists := elements[k-1]; exists {
			continue
		} else {
			length :=1 
			for {
			if _, ok := elements[k+length]; ok {
				length++
			} else {
				break
			}
			}

			if length > count {
				count = length
			}
		}
	}

	return count
}
