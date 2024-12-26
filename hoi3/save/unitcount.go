package save

type UnitCount struct {
	Count           int
	Type            string
	Strength        float64
	HighestStrength float64
	Organisation    float64
}

func (u *UnitCount) GetAverageOrganisation() float64 {
	return u.Organisation / float64(u.Count)
}

func (u *UnitCount) GetAverageStrength() float64 {
	return u.Strength / float64(u.Count)
}
