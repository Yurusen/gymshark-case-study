package service

import (
	"gymshark-case-study/internal/domain"
	"math"
	"sort"
)

type PackCalculator struct {
	packSizes []int
}

type state struct {
	totalPacks int
	counts     map[int]int
}

func NewPackCalculator(packSizes []int) *PackCalculator {
	sizes := make([]int, len(packSizes))
	copy(sizes, packSizes)
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))
	return &PackCalculator{packSizes: sizes}
}

func (pc *PackCalculator) CalculatePacks(items int) domain.PackResult {
	maxPack := pc.packSizes[0]
	limit := items + maxPack

	dp := make([]*state, limit+1)
	dp[0] = &state{totalPacks: 0, counts: map[int]int{}}

	for x := 0; x <= limit; x++ {
		if dp[x] == nil {
			continue
		}
		for _, size := range pc.packSizes {
			next := x + size
			if next > limit {
				continue
			}

			newCounts := make(map[int]int)
			for k, v := range dp[x].counts {
				newCounts[k] = v
			}
			newCounts[size]++

			newTotalPacks := dp[x].totalPacks + 1

			if dp[next] == nil ||
				(next >= items && (newTotalPacks < dp[next].totalPacks)) {
				dp[next] = &state{totalPacks: newTotalPacks, counts: newCounts}
			}
		}
	}

	best := map[int]int{}
	minOverage := math.MaxInt
	minPacks := math.MaxInt

	for x := items; x <= limit; x++ {
		if dp[x] != nil {
			overage := x - items
			if overage < minOverage || (overage == minOverage && dp[x].totalPacks < minPacks) {
				minOverage = overage
				minPacks = dp[x].totalPacks
				best = dp[x].counts
			}
		}
	}

	return domain.PackResult{Packs: best}
}
