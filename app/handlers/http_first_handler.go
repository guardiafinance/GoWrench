package handlers

import (
	contexts "wrench/app/contexts"
)

type HttpFirstHandler struct {
	Next Handler
}

func (httpFirst *HttpFirstHandler) Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	if httpFirst.Next != nil {
		httpFirst.Next.Do(wrenchContext, bodyContext)
	}
}

func (httpFirst *HttpFirstHandler) SetNext(next Handler) {
	httpFirst.Next = next
}
