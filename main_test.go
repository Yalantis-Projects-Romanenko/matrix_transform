package main

import (
	"fmt"
	"githib.com/fdistorted/matrixtransformer/mtx"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)

func TestSpiralTransform(t *testing.T) {
	matrix := [][]int{
		{1, 2},
		{4, 3},
	}
	slice := []int{1, 2, 3, 4}
	assert(Equal(mtx.SpiralTransform(matrix), slice))

	matrix = [][]int{
		{1, 2, 3},
		{6, 5, 4},
	}
	slice = []int{1, 2, 3, 4, 5, 6}
	assert(Equal(mtx.SpiralTransform(matrix), slice))

	matrix = [][]int{
		{1, 2, 3, 4, 5, 6, 7},
		{24, 25, 26, 27, 28, 29, 8},
		{23, 40, 41, 42, 43, 30, 9},
		{22, 39, 48, 49, 44, 31, 10},
		{21, 38, 47, 46, 45, 32, 11},
		{20, 37, 36, 35, 34, 33, 12},
		{19, 18, 17, 16, 15, 14, 13},
	}
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8,
		9, 10, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 20, 21, 22, 23,
		24, 25, 26, 27, 28, 29, 30,
		31, 32, 33, 34, 35, 36, 37,
		38, 39, 40, 41, 42, 43, 44,
		45, 46, 47, 48, 49}

	assert(Equal(mtx.SpiralTransform(matrix), slice))
}

func BenchmarkSpiralTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mtx.SpiralTransform(createMatrix(b.N,b.N))
	}
}

func createMatrix(x,y int)[][]int{
	matrix := make([][]int, y)
	for i := range matrix {
		matrix[i] = make([]int, x)
	}

	for i:=0; i<y; i++ {
		for j:=0; j<x; j++ {
			matrix[i][j] = i+j
		}
	}
	return matrix
}
// Equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func assert(o bool) {
	if !o {
		fmt.Printf("\n%c[35m%s%c[0m\n\n", 27, __getRecentLine(), 27)
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
