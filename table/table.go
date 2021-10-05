package table

import (
	"strings"

	goutils "github.com/simonski/goutils"
)

type Align int

const (
	AlignTable Align = iota
	AlignLeft
	AlignRight
	AlignCenter
)

type Column struct {
	Position     int
	Title        string
	WidthPercent int
	Align        Align
}

func (column *Column) Width() int {
	t := goutils.NewTerminal()
	return t.Width() / 100 * column.WidthPercent
}

type Table struct {
	Align           Align
	Columns         []*Column
	SeparatorWidth  int
	SeparatorColumn string
	SeparatorLine   string
	Border          bool
}

func NewTable() *Table {
	return &Table{Align: AlignLeft, Border: true, SeparatorWidth: 1, SeparatorColumn: "|", SeparatorLine: "-"}
}

func (table *Table) Add(column *Column) {
	table.Columns = append(table.Columns, column)
}

func (table *Table) Line() string {
	t := goutils.NewTerminal()

	// if border
	// 	|col|col|col|col|col|
	// 	So a 5-col table has 5+2 column separator
	// else
	// 	col|col|col|col|col
	// 	So a 5-col table has 5-1 column separators adding to width

	line := strings.Repeat(table.SeparatorLine, t.Width())
	return line
}

func (table *Table) Header() string {
	line := ""
	for _, column := range table.Columns {
		line += (column.Title + strings.Repeat(" ", column.Width()-len(column.Title)))
	}
	return line
}
