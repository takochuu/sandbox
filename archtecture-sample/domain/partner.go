package domain

type Partner struct {
}

func NewPartner() Partner {
	return Partner{}
}

func (d *Partner) GetPartner(partnerID int64) (User, error) {
	return User{}, nil
}
