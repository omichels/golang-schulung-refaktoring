package salary

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmployee_CalculatePay_IncludesBonusWhenWorkingOvertime(t *testing.T) {
	// given
	e := NewEmployee("europe") // hourlyPay is 20$
	e.SetClockedHours(40)
	// when
	sum := e.CalculatePay()
	//verify
	assert.Equal(t, 850, sum, "40 hours is overtime, 40 * 20 + 50 makes 850$")

	// given
	c := NewEmployee("china") // hourlyPay is 10$
	c.SetClockedHours(40)
	// when
	sum = c.CalculatePay()
	//verify
	assert.Equal(t, 400, sum, "40 hours is no overtime in china")
}

func Test_EmployeeFactory(t *testing.T) {
	e1 := NewEmployee("europe")
	c1 := NewEmployee("china")
	e3 := NewEmployee("not-china-not-europe")
	// verify
	assert.Equal(t, "europe", e3.GetLocation(), "fallback to europe as default")
	assert.NotEqual(t, e1.GetId(), c1.GetId())
	assert.NotEqual(t, e1.GetId(), e3.GetId())
	assert.NotEqual(t, e3.GetId(), c1.GetId())
}
