package kata1

type Status struct {
	hasInsurance bool
	isTotaled    bool
	isDented     bool
	isBigDent    bool
}

func (s Status) HasInsurance() bool {
	return s.hasInsurance
}

func (s Status) IsTotaled() bool {
	return s.isTotaled
}

func (s Status) IsDented() bool {
	return s.isDented
}

func (s Status) IsBigDent() bool {
	return s.isBigDent
}
