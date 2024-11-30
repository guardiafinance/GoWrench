package handlers

import (
	"encoding/json"
	contexts "wrench/app/contexts"
)

type HttpLastHandler struct {
	Next Handler
}

func (httpLast *HttpLastHandler) Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	var w = *wrenchContext.ResponseWriter
	w.WriteHeader(bodyContext.HttpStatusCode)
	json.NewEncoder(w).Encode(bodyContext.Body)

	if httpLast.Next != nil {
		httpLast.Next.Do(wrenchContext, bodyContext)
	}
}

func (httpLast *HttpLastHandler) SetNext(next Handler) {
	httpLast.Next = next
}
