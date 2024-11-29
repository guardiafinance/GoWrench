package handlers

import (
	"net/http"
	contexts "wrench/app/contexts"
)

type FirstHttpHandler struct {
	hasNext bool
	next    Handler
}

func FirstHttp(w http.ResponseWriter, r *http.Request) {
	firstHttpHandler := new(FirstHttpHandler)
	clientHander := HttpRequestClientHandler{}
	firstHttpHandler.SetNext(clientHander)

	wrenchContext := new(contexts.WrenchContext)
	bodyContext := new(contexts.BodyContext)

	firstHttpHandler.Do(wrenchContext, bodyContext)
}

func (firstHttp FirstHttpHandler) Do(wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {
	if firstHttp.hasNext {
		firstHttp.next.Do(wrenchContext, bodyContext)
	}
}

func (firstHttp FirstHttpHandler) SetNext(next Handler) {
	firstHttp.next = next
	firstHttp.hasNext = true
}
