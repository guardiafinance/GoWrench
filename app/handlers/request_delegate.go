package handlers

import (
	"net/http"
	contexts "wrench/app/contexts"
	settings "wrench/app/manifest/api_settings"
)

type RequestDelegate struct {
	Endpoint *settings.EndpointSettings
}

func (request *RequestDelegate) FirstHttp(w http.ResponseWriter, r *http.Request) {
	httpFirstHandler := new(HttpFirstHandler)
	httpLastHandler := new(HttpLastHandler)
	clientHander := new(HttpRequestClientHandler)
	httpFirstHandler.SetNext(clientHander)
	clientHander.SetNext(httpLastHandler)

	bodyContext := new(contexts.BodyContext)
	wrenchContext := new(contexts.WrenchContext)
	wrenchContext.ResponseWriter = &w

	httpFirstHandler.Do(wrenchContext, bodyContext)
}

func (request *RequestDelegate) SetEndpoint(endpoint *settings.EndpointSettings) {
	request.Endpoint = endpoint
}
