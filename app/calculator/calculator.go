package calculator

import (
	"maps"
	"math"
)

type solution struct {
	minShippedItems   int
	minPacksNumber    int
	packConfiguration map[int]int
}

// CalculateOrderConfiguration uses Dynamic Programming to build interim optimal interim solutions for each order size upto Order Size + Max Pack overship.
// And then backtracks to pick the best solution that ensures min overshipped items in min amount of packs.
func CalculateOrderConfiguration(availablePackSizes []int, orderSize int) map[int]int {
	if len(availablePackSizes) < 1 || orderSize < 1 {
		return map[int]int{}
	}

	biggestPack := availablePackSizes[len(availablePackSizes)-1]
	solutionsSize := biggestPack + orderSize + 1

	solutions := make([]*solution, solutionsSize)
	solutions[0] = &solution{minShippedItems: 0, minPacksNumber: 0, packConfiguration: map[int]int{}}

	// build interim optimal solutions upto orderSize + max pack size overship
	for i := 1; i < solutionsSize; i++ {
		for _, packSize := range availablePackSizes {
			if i-packSize >= 0 && solutions[i-packSize] != nil {
				previousSolution := solutions[i-packSize]
				newMinItems, newMinPacks := previousSolution.minShippedItems+packSize, previousSolution.minPacksNumber+1

				if solutions[i] == nil || newMinItems < solutions[i].minShippedItems || (newMinItems == solutions[i].minShippedItems && newMinPacks < solutions[i].minPacksNumber) {
					packCfg := map[int]int{}
					maps.Copy(packCfg, previousSolution.packConfiguration)
					packCfg[packSize]++
					solutions[i] = &solution{
						minShippedItems:   newMinItems,
						minPacksNumber:    newMinPacks,
						packConfiguration: packCfg,
					}
				}
			}
		}
	}

	itemsOverSent := math.MaxInt - orderSize
	packsSent := orderSize/availablePackSizes[0] + 1
	result := map[int]int{}
	// Select min overshipped items with min packs sent
	for i := orderSize; i < solutionsSize; i++ {
		candidate := solutions[i]
		if candidate == nil {
			continue
		}
		currItemsOverSent := candidate.minShippedItems - orderSize
		if currItemsOverSent < itemsOverSent ||
			(currItemsOverSent == itemsOverSent && candidate.minPacksNumber < packsSent) {
			result = candidate.packConfiguration
			itemsOverSent = orderSize - candidate.minShippedItems
			packsSent = candidate.minPacksNumber
		}
	}

	return result
}
