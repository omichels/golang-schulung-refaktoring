package main

import (
	"fmt"
	"github.com/omichels/golang-schulung-refaktoring/pkg/salary"
)

func main() {
	e1 := salary.Employee{
		Id:           1,
		HourlyPay:    1,
		ClockedHours: 40,
		Bonus:        0,
	}
	ec1 := salary.Employee{
		Id:           2,
		HourlyPay:    1,
		ClockedHours: 38,
		Bonus:        100,
	}
	c1 := salary.ChinaEmployee{
		Emp: ec1,
	}
	employees := make([]salary.PayCalculator, 0)
	employees = append(employees, e1)
	employees = append(employees, c1)
	totalSum := CalculateTotalSum(employees)
	fmt.Println(totalSum)
}

func CalculateTotalSum(employees []salary.PayCalculator) (total int) {
	sum := 0
	for _, v := range employees {
		sum = sum + v.CalculatePay()
	}
	return sum
}
