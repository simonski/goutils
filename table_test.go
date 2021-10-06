package goutils

import (
	"fmt"
	"testing"
)

func TestTable(test *testing.T) {
	fmt.Printf("Hello, World!\n")

	t := NewTable()
	t.Add(&Column{Position: 0, Title: "Fred", WidthPercent: 10})
	t.Add(&Column{Position: 0, Title: "Jack", WidthPercent: 20})
	t.Add(&Column{Position: 0, Title: "Pete", WidthPercent: 30})
	t.Add(&Column{Position: 0, Title: "Jim", WidthPercent: 20})
	t.Add(&Column{Position: 0, Title: "Rob", WidthPercent: 10})
	t.Add(&Column{Position: 0, Title: "Finbar", WidthPercent: 10})
	// fmt.Printf("%v\n", t.Header())
}
