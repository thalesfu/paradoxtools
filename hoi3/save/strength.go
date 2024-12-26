package save

type UnitStrength struct {
	Strength        float64
	HighestStrength float64
	Organisation    float64
}

func (u *UnitStrength) GetAverageOrganisation() float64 {
	return u.Organisation / u.Strength
}
