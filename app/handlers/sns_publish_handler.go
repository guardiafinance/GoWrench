package handlers

import (
	"context"
	contexts "wrench/app/contexts"
	settings "wrench/app/manifest/action_settings"
)

type SnsPublishHandler struct {
	Next           Handler
	ActionSettings *settings.ActionSettings
}

func (handler *SnsPublishHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	if wrenchContext.HasError == false {
	}

	if handler.Next != nil {
		handler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (handler *SnsPublishHandler) SetNext(next Handler) {
	handler.Next = next
}
