package salary

type PayCalculator interface {
	CalculatePay() (sum int)
}

type Employee struct {
	Id           int
	HourlyPay    int
	ClockedHours int
	Bonus        int
}

func (e Employee) CalculatePay() (sum int) {
	if e.ClockedHours > 38 {
		return (e.HourlyPay * e.ClockedHours) + e.Bonus
	} else {
		return e.HourlyPay * e.ClockedHours
	}
}

type ChinaEmployee struct {
	Emp Employee
}

func (e ChinaEmployee) CalculatePay() (sum int) {
	if e.Emp.ClockedHours > 100 {
		return (e.Emp.HourlyPay * e.Emp.ClockedHours) + e.Emp.Bonus
	} else {
		return e.Emp.HourlyPay * e.Emp.ClockedHours
	}
}
