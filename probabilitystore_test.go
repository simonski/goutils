package goutils

import (
	"fmt"
	"testing"
)

// createProbabilityStore helper method gives me a populated store I can
// run tests on
func createProbabilityStore(line string) *ProbabilityStore {
	c := createCLI("")
	values := c.SplitStringToFloats(line, ",")
	store := NewProbabilityStore(values)
	return store
}

func TestAllocation1M(t *testing.T) {
	line1 := "1,2,3,4,5"
	store := createProbabilityStore(line1)
	expect := 2
	actual := store.Indexof(4.5)

	if actual != expect {
		t.Errorf("bin should allocate to %d not %d\n", expect, actual)
	}

}

func TestAllocationWorks(t *testing.T) {
	line1 := "1,2,3,4,5"
	store := createProbabilityStore(line1)

	// range should be 0 .. 1
	expect := 0
	actual := store.Indexof(0.2)
	if actual != expect {
		t.Errorf("bin should allocate to %d not %d\n", expect, actual)
	}

	// range should be 1 .. 2
	expect = 1
	actual = store.Indexof(1.2)
	if actual != expect {
		t.Errorf("bin should allocate to %d not %d\n", expect, actual)
	}

	// range should be 3 .. 6
	expect = 2
	actual = store.Indexof(3.2)
	if actual != expect {
		t.Errorf("bind should allocate to %d not %d\n", expect, actual)
	}

	// range should be 6 .. 10
	expect = 3
	actual = store.Indexof(6.5)
	if actual != expect {
		t.Errorf("bind should allocate to %d not %d\n", expect, actual)
	}

	// range should be 10 .. 15
	expect = 4
	actual = store.Indexof(11.2)
	if actual != expect {
		t.Errorf("bind should allocate to %d not %d\n", expect, actual)
	}

	// range should be 15 .. 100
	expect = 5
	actual = store.Indexof(37)
	if actual != expect {
		t.Errorf("bind should allocate to %d not %d\n", expect, actual)
	}

}

func TestPerformance_ologn_WorstCase(t *testing.T) {
	line1 := "1,2,3,4,5,6,7,8,9"
	store := createProbabilityStore(line1)
	fmt.Printf("There are %d bins when the line is %s\n", store.Length(), line1)
	if store.Length() != 10 {
		t.Errorf("There should be 10 bins.\n")
	}

	// ok the search folds right, so just allocate to the minimum value
	actual := store.Indexof(0.1)
	expect := 0
	if actual != expect {
		t.Errorf("Bin allocation is incorrect (expect %d got %d).\n", expect, actual)
	}

	expect = 4
	if store.CallCount != expect {
		t.Errorf("Performance is incorrect (should be o log n)")
	}
}

func TestPerformance_ologn_v2_WorstCase(t *testing.T) {
	line1 := "1,2,3,4,5,6,7,8,9,10"
	store := createProbabilityStore(line1)

	// ok the search folds right, so just allocate to the minimum value
	actual := store.Search_o_log_n(0.1)
	expect := 0
	if actual != expect {
		t.Errorf("Bin allocation is incorrect (expect %d got %d).\n", expect, actual)
	}

	expect = 4
	if store.CallCount != expect {
		t.Errorf("Performance is incorrect (should be o log n)")
	}
}


// run some performance tests on varying N sizes and iteration sizes then determine the callcounts to verify
// we are in o(n) and o(log n) 
func TestPerformance(t *testing.T) {

	nlow := 100
	nhigh := 1000000
	nincrement := 1000
	iterations := 10000000

	for n:=nlow; n<=nhigh; n+=nincrement {
		report_on := RunPerformanceTest_on(n, iterations)
		report_ologn := RunPerformanceTest_ologn(n, iterations)
	}

}