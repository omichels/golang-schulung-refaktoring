package main

import "fmt"

type Employee struct {
	id           int
	hourlyPay    int
	clockedHours int
}

func main() {
	employees := []Employee{
		{
			id:           1,
			hourlyPay:    20,
			clockedHours: 40,
		},
		{
			id:           2,
			hourlyPay:    23,
			clockedHours: 37,
		},
	}
	totalSum := CalculateTotalSum(employees)

	fmt.Println(totalSum)
}

func CalculateTotalSum(employees []Employee) (total int) {
	sum := 0
	for _, v := range employees {
		sum = sum + (v.hourlyPay * v.clockedHours)
	}
	return sum
}
