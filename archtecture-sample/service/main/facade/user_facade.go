package facade

import "github.com/takochuu/anaconda"

type LikeFacade struct {
}

func NewLikeFacade() LikeFacade {
	return LikeFacade{}
}

type PostLikeArgument struct {
	Me User
	Limit  int
	Offset int
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

	domain.NewLikeDomain()

	domain.DoLike()

	ret := NewGetUserListRetVal()
	ret.Result := users
	return err, nil
}
