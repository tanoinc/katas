package kata1

import "testing"

func TestGetInsuranceAmount(t *testing.T) {
	tests := []struct {
		status   Status
		expected int
	}{
		{
			Status{hasInsurance: false},
			1,
		},
		{
			Status{hasInsurance: true, isTotaled: true},
			10000,
		},
		{
			Status{hasInsurance: true, isDented: true},
			160,
		},
		{
			Status{hasInsurance: true, isDented: true, isBigDent: true},
			270,
		},
		{
			Status{hasInsurance: true},
			0,
		},
	}

	for _, test := range tests {
		result := GetInsuranceAmount(test.status)

		if result != test.expected {
			t.Errorf("For status %+v, expected %d, but got %d", test.status, test.expected, result)
		}
	}
}
