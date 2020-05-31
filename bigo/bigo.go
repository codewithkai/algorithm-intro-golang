package main

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type Algorithm func(int) (string, float64)

var all = []Algorithm{
	//O(1)
	func(n int) (string, float64) {
		var count float64
		count++
		return "O(1)", count
	},

	//O(log(n))
	func(n int) (string, float64) {
		var count float64
		for i := n; i >= 1; i = i / 2 {
			count++
		}
		return "O(log(n))", count
	},

	//O(n)
	func(n int) (string, float64) {
		var count float64
		for i := 1; i <= n; i++ {
			count++
		}
		return "O(n)", count
	},

	//O(nlog(n))
	func(n int) (string, float64) {
		var count float64
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j = j * 2 {
				count++
			}
		}
		return "O(nlog(n))", count
	},

	//O(n^2)
	func(n int) (string, float64) {
		var count float64
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				count++
			}
		}
		return "O(n^2)", count
	},

	//O(n^3)
	func(n int) (string, float64) {
		var count float64
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				for k := 1; k <= n; k++ {
					count++
				}
			}
		}
		return "O(n^3)", count
	},

	//O(2^n)
	//func(n int) (string, float64) {
	//	var count float64
	//	for i := 1.0; i <= math.Pow(2, float64(n)); i++ {
	//		count++
	//	}
	//	return "O(2^n)", count
	//},
}

func main() {
	n := 1000
	fmt.Printf("Input size: %d\n", n)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Algorithm\tCount\tTime")
	for _, f := range all {
		now := time.Now()
		name, count := f(n)
		duration := time.Since(now)
		fmt.Fprintf(w, "%s\t%g\t%v\n", name, count, duration)
	}
	w.Flush()
}
