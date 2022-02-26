package salary

import (
	internal "github.com/omichels/golang-schulung-refaktoring/internal"
	"github.com/sirupsen/logrus"
	"reflect"
)

type PayCalculator interface {
	CalculatePay() (sum int)
	SetClockedHours(int)
}

type Employee struct {
	Id           int
	hourlyPay    int
	clockedHours int
	bonus        int
}

func (e *Employee) GetId() int {
	return e.Id
}

func (e *Employee) CalculatePay() (sum int) {
	if e.clockedHours > internal.OvertimeHourReachedForBonus {
		return (e.hourlyPay * e.clockedHours) + e.bonus
	} else {
		return e.hourlyPay * e.clockedHours
	}
}

func (e *Employee) SetClockedHours(i int) {
	e.clockedHours = i
}

type ChinaEmployee struct {
	emp Employee
}

func (e *ChinaEmployee) CalculatePay() (sum int) {
	if e.emp.clockedHours > internal.OvertimeHourReachedForBonusInChina {
		return (e.emp.hourlyPay * e.emp.clockedHours) + e.emp.bonus
	} else {
		return e.emp.hourlyPay * e.emp.clockedHours
	}
}
func (e *ChinaEmployee) SetClockedHours(i int) {
	e.emp.clockedHours = i
}

// EuropeType factory
func InitializeEuropeType(id int) PayCalculator {
	return &Employee{
		Id:        id,
		hourlyPay: 20,
		bonus:     50,
	}
}

// ChinaType factory
func InitializeChinaType(id int) PayCalculator {
	return &ChinaEmployee{
		emp: Employee{
			Id:        id,
			hourlyPay: 10,
			bonus:     100,
		},
	}
}

func NewEmployee(location string, id int) PayCalculator {
	// Map of strings to factories
	classes := map[string]func(int) PayCalculator{
		"europe": InitializeEuropeType,
		"china":  InitializeChinaType,
	}
	keys := reflect.ValueOf(classes).MapKeys()
	validateInputsNewEmployeeFactory(&location, keys)
	newEmp := classes[location](id)
	return newEmp
}

func validateInputsNewEmployeeFactory(location *string, keys []reflect.Value) {
	found := false
	for _, v := range keys {
		if v.String() == *location {
			found = true
		}
	}
	if !found {
		logrus.Debug("fallback to default europe")
		*location = "europe"
	}
}
