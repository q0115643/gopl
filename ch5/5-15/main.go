package main

import (
	"fmt"
)

func min(vals ...int) (rst int) {
	if len(vals) == 0 {
		rst = 0
		return
	}
	rst = vals[0]
	for _, val := range vals {
		if rst > val {
			rst = val
		}
	}
	return
}

func max(vals ...int) (rst int) {
	if len(vals) == 0 {
		rst = 0
		return
	}
	rst = vals[0]
	for _, val := range vals {
		if rst < val {
			rst = val
		}
	}
	return
}

func min2(first int, vals ...int) (rst int) {
	rst = first
	for _, val := range vals {
		if rst > val {
			rst = val
		}
	}
	return
}

func max2(first int, vals ...int) (rst int) {
	rst = first
	for _, val := range vals {
		if rst < val {
			rst = val
		}
	}
	return
}

func main() {
	fmt.Println(min(3, 4, 5))
	fmt.Println(max(3, 4, 5))
	fmt.Println(min2(3, 4, 5))
	fmt.Println(max2(3, 4, 5))
}
