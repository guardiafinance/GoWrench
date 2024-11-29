package handlers

import (
	contexts "wrench/app/contexts"
)

type Handler interface {
	Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext)
	SetNext(Handler)
}
