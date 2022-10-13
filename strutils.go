package goutils

import (
	"os"
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
	if times > 0 {
		return LPad(original, padding, times)
	}
	return original
}

func RPadToFixedLength(original string, padding string, maxLength int) string {
	times := maxLength - len(original)
	if times > 0 {
		return RPad(original, padding, times)
	}
	return original
}

func TokenswitchEnvironmentVariables(original string) string {
	home, _ := os.UserHomeDir()
	new_string := strings.ReplaceAll(original, "~", home)
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		key := "$" + pair[0]
		value := pair[1]
		new_string = strings.ReplaceAll(new_string, key, value)
	}
	return new_string
}

func StripWhitespace(line string) string {
	line = strings.ReplaceAll(line, " ", "")
	line = strings.ReplaceAll(line, "\t", "")
	line = strings.ReplaceAll(line, "\n", "")
	return line
}
