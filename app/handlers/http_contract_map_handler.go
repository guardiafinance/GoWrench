package handlers

import (
	"context"
	contexts "wrench/app/contexts"
)

type HttpContractMapHandler struct {
	Next Handler
}

func (httpFirst *HttpContractMapHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	if httpFirst.Next != nil {
		httpFirst.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (httpFirst *HttpContractMapHandler) SetNext(next Handler) {
	httpFirst.Next = next
}
