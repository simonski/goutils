package goutils

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err.Error())
		panic(err)
	}
}

type FileScanner struct {
	io.Closer
	*bufio.Scanner
}

func BuildScanner(filename *string) *FileScanner {
	if strings.HasSuffix(*filename, ".gz") {
		file, err := os.OpenFile(*filename, os.O_RDONLY, os.ModePerm)
		Check(err)
		gz, _ := gzip.NewReader(file)
		scanner := bufio.NewScanner(gz)
		return &FileScanner{file, scanner}
	} else {
		file, err := os.OpenFile(*filename, os.O_RDONLY, os.ModePerm)
		Check(err)
		scanner := bufio.NewScanner(file)
		return &FileScanner{file, scanner}
	}

}

func Console(msg ...string) {
	if len(msg) == 2 {
		fmt.Printf("%-30v%s\n", msg[0], msg[1])
	} else {
		fmt.Println(msg[0])
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Contains(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

var letterRunes = []rune(" -_.;:/1234567890)(*&^%$£abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Load_file_to_ints(filename string) []int {
	file, err := os.Open(filename)
	results := make([]int, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value, _ := strconv.Atoi(line)
		results = append(results, value)
	}
	return results

}

func Load_file_to_strings(filename string) []string {
	file, err := os.Open(filename)
	results := make([]string, 0)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results = append(results, line)
	}
	return results

}

// make_map_of_inty_list helper makes a map[int]int of a []int to give me
// whatever go's maps key algorithm performance is, at the cost of the memory
func Make_map_of_inty_list(data []int) map[int]int {
	m := make(map[int]int)
	for index := 0; index < len(data); index++ {
		value := data[index]
		m[value] = value
	}
	return m
}

func Convert_strings_to_ints(input []string) []int {
	output := make([]int, 0)
	for _, value := range input {
		ivalue, _ := strconv.Atoi(value)
		output = append(output, ivalue)
	}
	return output
}

func Min(v1 int, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func Max(v1 int, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}

func Decimal_to_binary(value int64) string {
	b := NewBitSet(int64(value))
	return b.ToBinaryString(36)
}

func Binary_to_decimal(decimalValue string) int64 {
	total := int64(0)
	for index := 0; index < len(decimalValue); index++ {
		value := decimalValue[index : index+1]
		power := len(decimalValue) - index - 1
		if value == "1" {
			total += int64(math.Pow(2, float64(power)))
		}
	}
	return total
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// FileExists indicates if a file already exists... or not
func FileExists(filename string) bool {
	result, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !result.IsDir()
}

// DirExists indicates if a file already exists... or not
func DirExists(filename string) bool {
	result, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return result.IsDir()
}

// EvaluateFilename replaces ~ with os.Getenv("HOME") on the filename
func EvaluateFilename(filename string) string {
	home := os.Getenv("HOME")
	newname := strings.ReplaceAll(filename, "~", home)
	return newname
}

func Intify(value string) int {
	ival, _ := strconv.Atoi(strings.TrimSpace(value))
	return ival
}

func Isint(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func Bitwisenot(value int) int {
	result := (1 << 16) - 1 - value
	// fmt.Printf("bitwisenot: ^%v=%v\n", value, result)
	return result
}

func Repeatstring(s string, times int) string {
	out := s
	for index := 0; index < times; index++ {
		out += s
	}
	return out
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Min64(a int64, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max64(a int64, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func MinU64(a uint64, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}

func MaxU64(a uint64, b uint64) uint64 {
	if a > b {
		return a
	}
	return b
}

func Factorial(a uint64) uint64 {
	if a > 1 {
		a = a * Factorial(a-1)
		return a
	} else {
		return a
	}
}

// https://en.wikipedia.org/wiki/Arithmetic_progression
// also: https://www.youtube.com/watch?v=uACt9OntiLo
func ArithmeticProgression(first int, last int) int {

	// number N terms being added (here, 5)
	// multiplying the sum of the first and last number then divide by 2
	//
	return (last * (first + last)) / 2

}
