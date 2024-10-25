package kata2

type Employee struct {
	rate      int
	separated bool
	retired   bool
}

func NewEmployee(rate int, separated, retired bool) Employee {
	return Employee{
		rate:      rate,
		separated: separated,
		retired:   retired,
	}
}

func (e Employee) GetRate() int {
	return e.rate
}

func (e Employee) IsSeparated() bool {
	return e.separated
}

func (e Employee) IsRetired() bool {
	return e.retired
}
