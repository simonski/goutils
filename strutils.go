package goutils

import (
	"strings"
)

func LPad(original string, padding string, repeat int) string {
	return strings.Repeat(padding, repeat) + original
}

func RPad(original string, padding string, repeat int) string {
	return original + strings.Repeat(padding, repeat)
}

func LPadToFixedLength(original string, padding string, maxLength int) string {
	times := maxLength - len(original)
	return LPad(original, padding, times)
}

func RPadToFixedLength(original string, padding string, maxLength int) string {
	times := maxLength - len(original)
	return RPad(original, padding, times)
}
