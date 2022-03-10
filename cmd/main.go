package main

import (
	"fmt"
	"github.com/omichels/golang-schulung-refaktoring/pkg/salary"
)

func main() {

	e1 := salary.NewEmployee("europe")
	e1.SetClockedHours(40)
	c1 := salary.NewEmployee("china")
	c1.SetClockedHours(20)
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
