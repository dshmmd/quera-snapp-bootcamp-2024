package w1q3

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

const pricePerItem = 8000000

var discountRates = map[int]float64{
	2: 0.95,
	3: 0.90,
	4: 0.80,
	5: 0.75,
}

// calculateGroupPrice Calculates the price of a group
func calculateGroupPrice(group []string) int {
	basePrice := pricePerItem

	uniqueTypes := make(map[string]bool)
	for _, item := range group {
		uniqueTypes[item] = true
	}
	uniqueTypeCount := len(uniqueTypes)

	discount := 1.0
	if uniqueTypeCount > 1 && uniqueTypeCount <= 5 {
		discount = discountRates[uniqueTypeCount]
	}

	discountedPrice := int(float64(uniqueTypeCount*basePrice) * discount)

	nonDiscountedCount := len(group) - uniqueTypeCount
	nonDiscountedPrice := nonDiscountedCount * basePrice

	totalPrice := discountedPrice + nonDiscountedPrice

	return totalPrice
}

// calculateTotalCost Calculates the total cost of a given grouping configuration
func calculateTotalCost(groups [][]string) int {
	totalCost := 0
	for _, group := range groups {
		totalCost += calculateGroupPrice(group)
	}
	return totalCost
}

// generateAllPartitions recursively generates all possible partitions of the input items.
// Each partition is a slice of groups, and each group is a slice of strings.
func generateAllPartitions(items []string) [][][]string {
	if len(items) == 0 {
		return [][][]string{{}}
	}

	first := items[0]
	rest := items[1:]
	subPartitions := generateAllPartitions(rest)

	var allPartitions [][][]string

	for _, partition := range subPartitions {
		for i := 0; i < len(partition); i++ {
			newPartition := deepCopyPartition(partition)

			newPartition[i] = append(newPartition[i], first)
			allPartitions = append(allPartitions, newPartition)
		}

		newGroup := []string{first}
		newPartition := deepCopyPartition(partition)
		newPartition = append(newPartition, newGroup)
		allPartitions = append(allPartitions, newPartition)
	}

	return allPartitions
}

// deepCopyPartition creates a deep copy of a partition to avoid modifying the original.
func deepCopyPartition(partition [][]string) [][]string {
	newPartition := make([][]string, len(partition))
	for i, group := range partition {
		newGroup := make([]string, len(group))
		copy(newGroup, group)
		newPartition[i] = newGroup
	}
	return newPartition
}

func Solve(src io.Reader) (answer string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()

	var input string
	_, err = fmt.Fscan(src, &input)
	if err != nil {
		return
	}

	items := strings.Split(input, ",")

	allPartitions := generateAllPartitions(items)

	minPrice := math.MaxInt32
	for _, groups := range allPartitions {
		cost := calculateTotalCost(groups)
		if cost < minPrice {
			minPrice = cost
		}
	}

	answer += fmt.Sprintln(minPrice)
	return
}

func init() {
	// make sure "os" package is present for submission script
	_ = os.DevNull
}
