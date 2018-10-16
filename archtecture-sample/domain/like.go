package domain

type Like struct {
}

func NewLike() Like {
	return Like{}
}

func (d *Like) DoLike(me, partner User) err {

}
