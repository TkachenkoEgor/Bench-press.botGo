package main

import (
	"fmt"
	"math"
	_ "os"
)

func main() {
	var PM float64
	fmt.Scan(&PM)

	var firstweek float64 = 0.6
	var secondweek float64 = 0.7
	var thirdweek float64 = 0.8
	var fourthweek float64 = 0.9

	first := PM * firstweek
	second := PM * secondweek
	third := PM * thirdweek
	fourth := PM * fourthweek

	fmt.Println("1 и 2 неделя", (math.Floor(first*1000) / 1000), "кг 4 подхода 10 повторений")
	fmt.Println("3 и 4 неделя", (math.Floor(second*1000) / 1000), "кг 5 подходов 6 повторений")
	fmt.Println("5 и 6 неделя", (math.Floor(third*1000) / 1000), "кг 6 подходов 5 повторений")
	fmt.Println("7 и 8 неделя", (math.Floor(fourth*1000) / 1000), "кг 3 подхода на 3 повторения")
}
