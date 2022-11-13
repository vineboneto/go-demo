package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Family struct {
	Name string
	Age  int
}

func timeTrack(start time.Time) {
	fmt.Println("total: ", time.Since(start))
}

func reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func mapper[T, M any](s []T, f func(T) M) []M {
	var arr []M
	for _, v := range s {
		newValue := f(v)
		arr = append(arr, newValue)
	}
	return arr
}

func sortExamples() {
	s := []int{4, 2, 3, 1}
	sort.Ints(s)
	fmt.Println(s)

	family := []Family{
		{"00", 23},
		{"11", 2},
		{"111", 2},
		{"112", 25},
		{"001", 25},
	}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Name < family[j].Name
	})
	fmt.Println(family)
}

func reduceExamples() {

	content, _ := os.ReadFile("./numbers.csv")

	records := strings.Split(string(content), ",")

	numbers := mapper(records, func(e string) int {
		v, _ := strconv.Atoi(e)
		return v
	})

	sum := reduce(numbers, func(acc int, current int) int {
		return acc + current
	}, 0)
	fmt.Println(sum)
}

func main() {
	defer timeTrack(time.Now())
	reduceExamples()
}
