package controller

import (
	"github.com/eure/eure-payment/library/net/http/response"
	"github.com/golang/glog"
	"github.com/takochuu/sandbox/archtecture-sample/service/main/facade"
)

type LikeController struct {
}

func NewLikeController() LikeController {
	return LikeController{}
}

func (c *LikeController) PostLike(r request) {
	// Transactionのスタート
	transaction.Begin()

	// Facade
	f := facade.NewLikeFacade()
	opt := facade.NewPostLikeArgument()
	ret, err := f.PostLike(opt)
	if err != nil {
		// Rollback
		transaction.Rollback()

		// Logging
		glog.Errorln(err)

		// Response
		response.ErrorJSON()
	}

	// Transactionのコミット
	transaction.Commit()

	// レスポンスの集計Converter
	converter.NewPostLike()
}
