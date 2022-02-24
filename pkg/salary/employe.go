package salary

type Employee struct {
	Id           int
	HourlyPay    int
	ClockedHours int
	Bonus        int
}

func (e Employee) CalculatePay() (sum int) {
	if e.ClockedHours > 38 {
		return e.HourlyPay*e.ClockedHours + e.Bonus
	} else {
		return e.HourlyPay * e.ClockedHours
	}
}
