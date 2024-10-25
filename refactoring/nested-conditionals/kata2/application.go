package kata2

func PayAmount(employee Employee, workHours int) PayCheck {
	if employee.IsSeparated() {
		return NewPayCheck(0, "SEP")

	}
	if employee.IsRetired() {
		return NewPayCheck(0, "RET")
	}

	bonus := computeBonus(workHours)
	regularAmount := computeRegularPayAmount(employee, workHours)
	amount := bonus + regularAmount

	return NewPayCheck(amount, "EMP")

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
