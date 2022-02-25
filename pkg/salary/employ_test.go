package salary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployee_CalculatePay_IncludesBonusWhenWorkingOvertime(t *testing.T) {
	// given
	e := Employee{
		Id:           1,
		HourlyPay:    1,
		ClockedHours: 40,
		Bonus:        10,
	}
	// when
	sum := e.CalculatePay()
	//verify
	assert.Equal(t, 50, sum, "40 hours is overtime, 40 * 1 + 10 makes 50$")

	// given
	c := ChinaEmployee{}
	c.Emp = e
	// when
	sum = c.CalculatePay()
	//verify
	assert.Equal(t, 40, sum, "40 hours is no overtime in china")
}
