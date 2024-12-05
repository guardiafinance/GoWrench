package handlers

import (
	"context"
	"strings"
	client "wrench/app/clients/http"
	"wrench/app/contexts"
	settings "wrench/app/manifest/action_settings"
)

type HttpRequestClientHandler struct {
	Next           Handler
	ActionSettings *settings.ActionSettings
}

func (httpRequestClientHandler *HttpRequestClientHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	if wrenchContext.HasError == false {

		request := new(client.HttpClientRequestData)
		request.Body = []byte(bodyContext.Body)
		request.Method = string(httpRequestClientHandler.ActionSettings.Request.Method)
		request.Url = httpRequestClientHandler.ActionSettings.Request.Url

		response, err := client.HttpClientDo(ctx, request)

		if err != nil {
			wrenchContext.SetHasError()
		}

		if response.StatusCode > 399 {
			wrenchContext.SetHasError()
		}

		bodyContext.Body = string(response.Body)
		bodyContext.HttpStatusCode = response.StatusCode
		bodyContext.SetHeaders(httpRequestClientHandler.ActionSettings.Request.MapFixedHeaders)
		bodyContext.SetHeaders(mpaHttpResponseHeaders(response, httpRequestClientHandler.ActionSettings.Request.MapResponseHeader))
	}

	if httpRequestClientHandler.Next != nil {
		httpRequestClientHandler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (httpRequestClientHandler *HttpRequestClientHandler) SetNext(handler Handler) {
	httpRequestClientHandler.Next = handler
}

func mpaHttpResponseHeaders(response *client.HttpClientResponseData, mapResponseHeader []string) map[string]string {

	if mapResponseHeader == nil {
		return nil
	}
	mapResponseHeaderResult := make(map[string]string)

	for _, mapHeader := range mapResponseHeader {
		mapSplitted := strings.Split(mapHeader, ":")
		sourceKey := mapSplitted[0]
		var destinationKey string
		if len(mapSplitted) > 1 {
			destinationKey = mapSplitted[1]
		}

		if len(destinationKey) == 0 {
			destinationKey = sourceKey
		}

		headerValue := response.HttpClientResponse.Header.Get(sourceKey)
		mapResponseHeaderResult[destinationKey] = headerValue
	}

	return mapResponseHeaderResult
}
