package kata1

func GetInsuranceAmount(status Status) int {
	if !status.HasInsurance() {
		return 1
	}

	if status.IsTotaled() {
		return 10000
	}

	if !status.IsDented() {
		return 0
	}

	if status.IsBigDent() {
		return 270
	}

	return 160
}
