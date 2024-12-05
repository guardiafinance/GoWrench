package handlers

import (
	"context"
	client "wrench/app/clients/http"
	"wrench/app/contexts"
	settings "wrench/app/manifest/action_settings"
)

type HttpRequestClientHandler struct {
	Next           Handler
	ActionSettings *settings.ActionSettings
}

func (httpRequestClientHandler *HttpRequestClientHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	request := new(client.HttpClientRequestData)
	request.Body = []byte(bodyContext.Body)
	request.Method = string(httpRequestClientHandler.ActionSettings.Request.Method)
	request.Url = httpRequestClientHandler.ActionSettings.Request.Url

	response, err := client.HttpClientDo(ctx, request)

	if err != nil {
		//TODO add breaking setting error status
	}

	bodyContext.Body = string(response.Body)
	bodyContext.HttpStatusCode = response.StatusCode
	bodyContext.SetHeaders(httpRequestClientHandler.ActionSettings.Request.FixedHeaders)

	if httpRequestClientHandler.Next != nil {
		httpRequestClientHandler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (httpRequestClientHandler *HttpRequestClientHandler) SetNext(handler Handler) {
	httpRequestClientHandler.Next = handler
}
