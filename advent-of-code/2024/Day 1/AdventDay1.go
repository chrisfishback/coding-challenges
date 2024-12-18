package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lines := 0
	var left [1000]int
	var right [1000]int

	data, err := os.Open("./data")
	check(err)
	defer data.Close()

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")
		// fmt.Println(numbers[3])
		left[lines], err = strconv.Atoi(numbers[0])
		check(err)
		right[lines], err = strconv.Atoi(numbers[3])
		check(err)
		lines++
	}

	leftNums := left[:lines]
	sort.Ints(leftNums)
	rightNums := right[:lines]
	sort.Ints(rightNums)

	var sum = 0.0

	var i int
	for i = 0; i < len(leftNums); i++ {
		//fmt.Println(leftNums[i], " - ", rightNums[i], " = ", math.Abs(float64(leftNums[i]-rightNums[i])))
		fmt.Println(sum)
		sum = sum + math.Abs(float64(leftNums[i]-rightNums[i]))
	}

	fmt.Println(sum)
}
