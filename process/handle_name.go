package process

import (
	"context"
	"strings"

	"github.com/guoruibiao/responsibility-chain/params"
)

func UpperName(ctx *context.Context, reqParams *params.CommonParams)(res interface{}, err error) {
	reqParams.Name = strings.ToUpper(reqParams.Name)
	*ctx = context.WithValue(*ctx, "name", reqParams.Name)
	return
}