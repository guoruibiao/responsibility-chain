package process

import (
	"context"

	"github.com/guoruibiao/responsibility-chain/params"
)

func ModifyAddress(ctx *context.Context, reqParams *params.CommonParams)(res interface{}, err error) {
	switch reqParams.Address {
	case "北京":
		reqParams.Address = "Beijing"
	case "南京":
		reqParams.Address = "Nanjing"
	case "大连":
		reqParams.Address = "Dalian"
	}

	*ctx = context.WithValue(*ctx, "address", reqParams.Address)
	return
}