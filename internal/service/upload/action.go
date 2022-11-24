package upload

import (
	"context"
	"github.com/solost23/go_interface/gen_go/oss"
	"oss_service/internal/service/base"
)

type Action struct {
	base.Action
}

func NewActionWithCtx(ctx context.Context) *Action {
	a := &Action{}
	a.SetContext(ctx)
	return a
}

func (a *Action) Deal(ctx context.Context, request *oss.UploadRequest) (reply *oss.UploadResponse, err error) {
	return nil, nil
}
