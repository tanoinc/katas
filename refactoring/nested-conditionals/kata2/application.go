package kata2

func PayAmount(employee Employee, workHours int) PayCheck {
	var result PayCheck
	if !employee.IsSeparated() {
		if employee.IsRetired() {
			result = NewPayCheck(0, "RET")
		} else {
			bonus := computeBonus(workHours)
			regularAmount := computeRegularPayAmount(employee, workHours)
			amount := bonus + regularAmount
			result = NewPayCheck(amount, "EMP")
		}
	} else {
		result = NewPayCheck(0, "SEP")
	}
	return result
}

func computeBonus(workHours int) int {
	if workHours > 40 {
		return 1000
	}
	return 0
}

func computeRegularPayAmount(employee Employee, workHours int) int {
	return employee.GetRate() * workHours
}
