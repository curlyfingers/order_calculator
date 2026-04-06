package calculator

import (
	"reflect"
	"testing"
)

var (
	packSizes = []int{250, 500, 1000, 2000, 5000}
)

func TestCalculateOrderConfiguration(t *testing.T) {
	testCases := []struct {
		name      string
		orderSize int
		expected  map[int]int
	}{
		{
			name:      "Zero items",
			orderSize: 0,
			expected:  map[int]int{},
		},
		{
			name:      "Sub smallest pack items",
			orderSize: packSizes[0] - 1,
			expected:  map[int]int{packSizes[0]: 1},
		},
		{
			name:      "Exactly smallest pack items",
			orderSize: packSizes[0],
			expected:  map[int]int{packSizes[0]: 1},
		},
		{
			name:      "Over smallest pack, but under next pack items",
			orderSize: packSizes[0] + packSizes[1]/2 - 1,
			expected:  map[int]int{packSizes[1]: 1},
		},
		{
			name:      "Smallest + next - 10",
			orderSize: packSizes[0] + packSizes[1] - 10,
			expected:  map[int]int{packSizes[1]: 1, packSizes[0]: 1},
		},
		{
			name:      "749",
			orderSize: 749,
			expected:  map[int]int{packSizes[0]: 1, packSizes[1]: 1},
		},
		{
			name:      "751",
			orderSize: 751,
			expected:  map[int]int{1000: 1},
		},
		{
			name:      "1750",
			orderSize: 1750,
			expected:  map[int]int{1000: 1, 500: 1, 250: 1},
		},
		{
			name:      "1751",
			orderSize: 1751,
			expected:  map[int]int{2000: 1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := CalculateOrderConfiguration(packSizes, tc.orderSize)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected %v, but got %v instead", tc.expected, actual)
			}
		})
	}
}

func TestCalculateOrderConfiguration_NoAvailablePackSizes(t *testing.T) {
	expected := map[int]int{}
	actual := CalculateOrderConfiguration([]int{}, 42)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v, but got %v instead", expected, actual)
	}
}
