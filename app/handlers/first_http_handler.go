package handlers

import (
	"net/http"
	contexts "wrench/app/contexts"
)

type FirstHttpHandler struct {
	next Handler
}

func FirstHttp(w http.ResponseWriter, r *http.Request) {
	firstHttpHandler := new(FirstHttpHandler)

	wrenchContext := new(contexts.WrenchContext)
	bodyContext := new(contexts.BodyContext)
	firstHttpHandler.Do(wrenchContext, bodyContext)
}

func (firstHttp *FirstHttpHandler) Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	firstHttp.next.Do(wrenchContext, bodyContext)
}

func (firstHttp *FirstHttpHandler) SetNext(next Handler) {
	firstHttp.next = next
}
