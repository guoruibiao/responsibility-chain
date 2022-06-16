package process

import (
	"context"

	"github.com/guoruibiao/responsibility-chain/params"
)

func DecreaseAge(ctx *context.Context, reqParams *params.CommonParams)(res interface{}, err error) {
	reqParams.Age = reqParams.Age - 1
	*ctx = context.WithValue(*ctx, "age", reqParams.Age)
	return
}
