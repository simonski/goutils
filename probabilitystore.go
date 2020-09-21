package goutils

import (
	"fmt"
	)

// Bin helper to hold the proability and nubmer of times it has occured
type Bin struct {
	Index 		int
	Probability float64
	Count       int
	LowerBound 	float64
	UpperBound  float64
}

type ProbabilityStore struct {
	bins []*Bin
	CallCount int
	Misses int
}

// Debug prints the results to stdout
func (bins ProbabilityStore) Debug (totalRows int) {
	fmt.Println("")
	pct := float64(100) / float64(totalRows)
	for index := 0; index < bins.Length(); index++ {
		bin := bins.bins[index]
		binPct := pct * float64(bin.Count)
		difference := bin.Probability - binPct
		fmt.Printf("Bin[%d] requested %.2f pct, (lower %.2f/upper %.2f), received %d hits, achieved %.3f pct, difference %.3f pct\n", 
			index, bin.Probability, bin.LowerBound, bin.UpperBound, bin.Count, binPct, difference)
	}
	fmt.Printf("BS.CallCount %d\n", bins.CallCount)
	fmt.Printf("BS.MissCount %d\n", bins.Misses)
	fmt.Println("")
}


func (b ProbabilityStore) Length () int {
	return len(b.bins)
}

// IndexOf return the position in the array of the Bin that 
// serves the value or -1 if it does not exist
func (b *ProbabilityStore) Indexof (value float64) int {
	index := b.BinarySearch(value, b.bins)
	return index
}

func (b *ProbabilityStore) BinarySearch(value float64, bins []*Bin) int {

	b.CallCount += 1

	if len(bins) == 0 {
		return -1
	}

	index := len(bins)/2
	candidate := bins[index]

	if candidate.UpperBound > value && candidate.LowerBound < value {	// eq
		// this is it
		return candidate.Index
	} else if candidate.UpperBound < value {					// lt
		// drop all to the left
		searchSpace := bins[index:]
		return b.BinarySearch(value, searchSpace)
	} else if candidate.LowerBound > value {					// gt
		// drop all to the right
		searchSpace := bins[0:index]
		return b.BinarySearch(value, searchSpace)
	} else {
		// we don't have it
		b.Misses += 1
		return -1
	}

}


func NewProbabilityStore(values[] float64) *ProbabilityStore {
	bins := make([]*Bin, len(values))
	lower := 0.0
	upper := 0.0
	remainder := float64(100)
	for index := 0; index<len(values); index++ {
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
		bin := Bin{lastBin.Index+1, remainder, 0, lastBin.UpperBound, 100.0}
		bins = append(bins, &bin)
	}
	bs := ProbabilityStore{bins, 0, 0}
	return &bs
}