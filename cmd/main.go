package main

import (
	"fmt"
	"github.com/omichels/golang-schulung-refaktoring/pkg/salary"
)

func main() {
	employees := []salary.Employee{
		{
			Id:           1,
			HourlyPay:    1,
			ClockedHours: 40,
			Bonus:        10,
		},
		{
			Id:           2,
			HourlyPay:    1,
			ClockedHours: 37,
		},
	}
	totalSum := CalculateTotalSum(employees)
	fmt.Println(totalSum)
}

func CalculateTotalSum(employees []salary.Employee) (total int) {
	sum := 0
	for _, v := range employees {
		sum = sum + v.CalculatePay()
	}
	return sum
}
