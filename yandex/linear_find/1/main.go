package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scanf("%d", &num)

	var minX = -1e9
	var maxX = 1e9

	var minY = -1e9
	var maxY = 1e9

	for j := 0; j < num; j++ {
		var x, y float64
		fmt.Scanf("%f %f", &x, &y)

		if x > minX {
			minX = x
		}

		if y < maxY {
			maxY = y
		}

		if x < maxX {
			maxX = x
		}

		if y > minY {
			minY = y
		}
	}

	fmt.Println(maxX, maxY, minX, minY)
}
