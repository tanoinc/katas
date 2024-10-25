package kata2

type PayCheck struct {
	amount int
	status string
}

func NewPayCheck(amount int, status string) PayCheck {
	return PayCheck{
		amount: amount,
		status: status,
	}
}
