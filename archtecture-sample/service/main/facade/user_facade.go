package facade

import "github.com/takochuu/sandbox/archtecture-sample/domain"

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
	p := domain.NewPartner()
	partner, err := p.GetPartner(opt.PartnerID)
	if err != nil {
		return ret, err
	}

	l := domain.NewLike()
	err = l.DoLike(opt.Me, partner)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
