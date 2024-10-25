package kata2

import "testing"

func TestPayAmount(t *testing.T) {
	tests := []struct {
		employee  Employee
		workHours int
		expected  PayCheck
	}{
		{
			NewEmployee(50, false, false),
			45,
			NewPayCheck(3250, "EMP"),
		},
		{
			NewEmployee(50, false, true),
			45,
			NewPayCheck(0, "RET"),
		},
		{
			NewEmployee(50, true, false),
			45,
			NewPayCheck(0, "SEP"),
		},
		{
			NewEmployee(50, false, false),
			35,
			NewPayCheck(1750, "EMP"),
		},
	}

	for _, test := range tests {
		result := PayAmount(test.employee, test.workHours)
		if result != test.expected {
			t.Errorf("For employee %+v and workHours %d, expected %+v, but got %+v", test.employee, test.workHours, test.expected, result)
		}
	}
}
