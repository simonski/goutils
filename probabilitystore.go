package goutils

import (
	"fmt"
)

// Bin helper to hold the probability and number of times it has occured
type Bin struct {
	Index       int     // position in original array
	Probability float64 // chance of this event occurring
	Count       int     // number of times it has occurred in this run
	LowerBound  float64 // lvalue in range of 0 .. 1 to accept
	UpperBound  float64 // rvalue in range of 0 .. 1 to accept
}

// ProbabilityStore helper struct holds all Bins and methods
type ProbabilityStore struct {
	Bins      []*Bin
	CallCount int
}

// Debug prints the results to stdout
func (store ProbabilityStore) Debug(totalRows int) {
	fmt.Println("")
	pct := float64(100) / float64(totalRows)
	for index := 0; index < store.Length(); index++ {
		bin := store.Bins[index]
		binPct := pct * float64(bin.Count)
		difference := bin.Probability - binPct
		fmt.Printf("Bin[%d] requested %.2f pct, (lower %.2f/upper %.2f), received %d hits, achieved %.3f pct, difference %.3f pct\n",
			index, bin.Probability, bin.LowerBound, bin.UpperBound, bin.Count, binPct, difference)
	}
	fmt.Printf("store.CallCount %d\n", store.CallCount)
	fmt.Println("")
}

// Length returns integer number of Bins in the store
func (store ProbabilityStore) Length() int {
	return len(store.Bins)
}

// Indexof return the position in the array of the Bin that
// serves the value or -1 if it does not exist
func (store *ProbabilityStore) Indexof(value float64) int {
	index := store.BinarySearch(value, store.Bins)
	return index
}

// BinarySearch performs a binary serach to find the appropriate bin for the passed value
func (store *ProbabilityStore) BinarySearch(value float64, bins []*Bin) int {

	store.CallCount++

	if len(bins) == 0 {
		return -1
	}

	index := len(bins) / 2
	candidate := bins[index]

	if candidate.UpperBound > value && candidate.LowerBound < value { // eq
		// this is it
		return candidate.Index
	} else if candidate.UpperBound < value { // lt
		// drop all to the left
		searchSpace := bins[index:]
		return store.BinarySearch(value, searchSpace)
	} else if candidate.LowerBound > value { // gt
		// drop all to the right
		searchSpace := bins[0:index]
		return store.BinarySearch(value, searchSpace)
	} else {
		// we don't have it
		return -1
	}

}

// Search_on performs a search in o(n) time to find the appropriate bin for the passed value
func (store *ProbabilityStore) Search_on(value float64) int {
	for index := 0; index < len(store.Bins); index++ {
		store.CallCount++
		candidate := store.Bins[index]
		if candidate.UpperBound > value && candidate.LowerBound < value { // eq
			// this is it
			return candidate.Index
		}
	}
	// we don't have it
	return -1
}

// Search_o_log_n performs a search in o(n) time to find the appropriate bin for the passed value
func (store *ProbabilityStore) Search_o_log_n(value float64) int {

	lindex := 0
	rindex := len(store.Bins)

	for loop := 0; loop < len(store.Bins)/2; loop++ {
		store.CallCount++

		middle := lindex + ((rindex - lindex) / 2)
		entry := store.Bins[middle]
		if entry.LowerBound > value {
			// chop all from right
			rindex = middle
		} else if entry.UpperBound < value {
			// chop all from left
			lindex = middle
		} else if entry.UpperBound > value && entry.LowerBound < value {
			// this is it
			return entry.Index
		}

	}
	return -1

}

// NewProbabilityStore factory function creates a store and allocates the Bins against teh array of probabilities
func NewProbabilityStore(values []float64) *ProbabilityStore {
	bins := make([]*Bin, len(values))
	lower := 0.0
	upper := 0.0
	remainder := float64(100)
	for index := 0; index < len(values); index++ {
		probability := values[index]
		lower = upper
		upper += probability
		if index == 0 {
			lower = 0
			upper = probability
		} else {
			lower = bins[index-1].UpperBound
			upper = lower + probability
		}

		bin := Bin{index, probability, 0, lower, upper}
		bins[index] = &bin
		remainder -= probability
	}
	if remainder > 0 {
		lastBin := bins[len(bins)-1]
		bin := Bin{lastBin.Index + 1, remainder, 0, lastBin.UpperBound, 100.0}
		bins = append(bins, &bin)
	}
	store := ProbabilityStore{bins, 0}
	return &store
}
