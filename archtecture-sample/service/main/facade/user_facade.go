package facade

type LikeFacade struct {
}

func NewLikeFacade() LikeFacade {
	return LikeFacade{}
}

type PostLikeArgument struct {
	Me        User
	PartnerID int64
	Limit     int
	Offset    int
}

func NewPostLikeArgument() PostLikeArgument {
	return PostLikeArgument{}
}

type PostLikeRetVal struct {
}

func NewGetUserListRetVal() PostLikeRetVal {
	return PostLikeRetVal{}
}

func (c *LikeFacade) PostLike(opt PostLikeArgument) (PostLikeRetVal, error) {
	ret := NewGetUserListRetVal()
	p := domain.NewPartnerDomain()
	partner, err := p.GetPartner(opt.PartnerID)
	if err != nil {
		return ret, err
	}

	l := domain.NewLikeDomain()
	err = l.DoLike(opt.Me, partner)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
