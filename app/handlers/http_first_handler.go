package handlers

import (
	"context"
	"io"
	contexts "wrench/app/contexts"
)

type HttpFirstHandler struct {
	Next Handler
}

func (httpFirst *HttpFirstHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	body, err := io.ReadAll(wrenchContext.Request.Body)

	if err != nil {
		wrenchContext.SetHasError()
	}

	bodyContext.BodyByteArray = body
	bodyContext.ContentType = "application/json"

	if httpFirst.Next != nil {
		httpFirst.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (httpFirst *HttpFirstHandler) SetNext(next Handler) {
	httpFirst.Next = next
}
