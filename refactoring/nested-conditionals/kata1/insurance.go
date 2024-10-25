package kata1

func GetInsuranceAmount(status Status) int {
	var amount int
	if !status.HasInsurance() {
		amount = 1
	} else {
		if status.IsTotaled() {
			amount = 10000
		} else {
			if status.IsDented() {
				amount = 160
				if status.IsBigDent() {
					amount = 270
				}
			} else {
				amount = 0
			}
		}
	}
	return amount
}
