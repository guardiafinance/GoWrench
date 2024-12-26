package handlers

import (
	"context"
	"fmt"
	"net/http"
	contexts "wrench/app/contexts"
	settings "wrench/app/manifest/api_settings"
	appSetting "wrench/app/manifest/application_settings"
)

type RequestDelegate struct {
	Endpoint *settings.EndpointSettings
}

func (request *RequestDelegate) HttpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	requestURI := r.RequestURI
	appSetting := appSetting.ApplicationSettingsStatic

	endpoint, err := appSetting.Api.GetEndpointByRoute(requestURI)
	if err != nil {
		fmt.Print(err)
	}

	var chain = ChainStatic.GetStatic()
	var handler = chain.GetByRoute(endpoint.Route)

	bodyContext := new(contexts.BodyContext)
	wrenchContext := new(contexts.WrenchContext)
	wrenchContext.ResponseWriter = &w
	wrenchContext.Request = r
	handler.Do(ctx, wrenchContext, bodyContext)
}

func (request *RequestDelegate) SetEndpoint(endpoint *settings.EndpointSettings) {
	request.Endpoint = endpoint
}
