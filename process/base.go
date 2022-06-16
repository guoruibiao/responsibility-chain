package process

import (
	"context"
	"sync"

	"github.com/guoruibiao/responsibility-chain/params"
)

type processor struct {
	bucket map[string]func(ctx *context.Context, reqParams *params.CommonParams)(interface{}, error)
	sortedChain []string
	sync.Mutex
}

var p processor

func Init() {
	 p = processor{
		 make(map[string]func(ctx *context.Context, reqParams *params.CommonParams)(interface{}, error)),
		 make([]string, 0),
		 sync.Mutex{},
	 }
}

func (p *processor)set(key string, callback func(ctx *context.Context, reqParams *params.CommonParams)(interface{}, error)) {
	p.Lock()
	defer p.Unlock()
	if _, exists := p.bucket[key]; !exists {
		p.bucket[key] = callback
		p.sortedChain = append(p.sortedChain, key)
	}
}

func Register(key string, callback func(ctx *context.Context, reqParams *params.CommonParams)(interface{}, error)) {
	p.set(key, callback)
}

func Run(ctx *context.Context, reqParams *params.CommonParams)(res interface{}, err error) {
	return p.run(ctx, reqParams)
}

func (p *processor)run(ctx *context.Context, reqParams *params.CommonParams) (res interface{}, err error) {
	for idx, _ := range p.sortedChain {
		res, err = p.bucket[p.sortedChain[idx]](ctx, reqParams)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

