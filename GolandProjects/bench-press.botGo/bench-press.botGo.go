package main

import (
	"fmt"
	_ "fmt"
	"math"
	"sort"
)

func main() {
	var PM float64
	fmt.Scan(&PM)
	coef := make(map[int]float64)
	var keys []int
	coef[1] = 0.6
	coef[2] = 0.7
	coef[3] = 0.8
	coef[4] = 0.9

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println("Key:", k, "Value:", coef[k])

	}

	for i, value := range coef {
		PMNEW := value * PM
		switch i {
		case 0:
			fmt.Println("начальный разовый максимум", PMNEW)

		case 1:
			fmt.Println("1 и 2 неделя", (math.Ceil(PMNEW)), "кг 4 подхода 10 повторений")
		case 2:
			fmt.Println("3 и 4 неделя", (math.Ceil(PMNEW)), "кг 5 подходов 6 повторений")
		case 3:
			fmt.Println("5 и 6 неделя", (math.Ceil(PMNEW)), "кг 6 подходов 5 повторений")
		case 4:
			fmt.Println("7 и 8 неделя", (math.Ceil(PMNEW)), "кг 3 подхода на 3 повторения")

		}

	}

}
