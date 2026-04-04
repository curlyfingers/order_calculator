package calculator

func CalculateOrderConfiguration(availablePackSizes []int, orderSize int) map[int]int {
	result := map[int]int{}

	if orderSize <= 0 || len(availablePackSizes) == 0 {
		return result
	}

	remainingItems := orderSize
	for i := len(availablePackSizes) - 1; i >= 0; i-- {
		wholePacks := remainingItems / availablePackSizes[i]
		result[availablePackSizes[i]] = wholePacks
		remainingItems = remainingItems - wholePacks*availablePackSizes[i]
	}

	if remainingItems > 0 {
		result[availablePackSizes[0]] = result[availablePackSizes[0]] + 1
	}

	for packSize, q := range result {
		if q == 0 {
			delete(result, packSize)
		}
	}

	return result
}
