package handlers

import (
	"log"
	action_settings "wrench/app/manifest/action_settings"
	settings "wrench/app/manifest/application_settings"
)

var ChainStatic *Chain = new(Chain)

type Chain struct {
	MapHandle map[string]Handler
}

func (chain *Chain) GetStatic() *Chain {
	return ChainStatic
}

func (chain *Chain) BuildChain(settings *settings.ApplicationSettings) {
	chain.MapHandle = make(map[string]Handler)
	for _, endpoint := range settings.Api.Endpoints {

		var firstHandler = new(HttpFirstHandler)
		var currentHandler Handler
		currentHandler = firstHandler

		var action, err = settings.GetActionById(endpoint.ActionID)

		if err != nil {
			log.Fatal(err)
		}

		if action.Trigger != nil && action.Trigger.Before != nil {
			httpContractMapHandler := new(HttpContractMapHandler)

			contractMapId := action.Trigger.Before.ContractMapId
			httpContractMapHandler.ContractMap = settings.Contract.GetContractById(contractMapId)

			currentHandler.SetNext(httpContractMapHandler)
			currentHandler = httpContractMapHandler
		}

		if action.Type == action_settings.ActionTypeHttpRequest {
			httpRequestHadler := new(HttpRequestClientHandler)
			httpRequestHadler.ActionSettings = action
			currentHandler.SetNext(httpRequestHadler)
			currentHandler = httpRequestHadler
		}

		if action.Type == action_settings.ActionTypeHttpRequestMock {
			httpRequestMockHadler := new(HttpRequestClientMockHandler)
			httpRequestMockHadler.ActionSettings = action
			currentHandler.SetNext(httpRequestMockHadler)
			currentHandler = httpRequestMockHadler
		}

		if action.Trigger != nil && action.Trigger.After != nil {
			httpContractMapHandler := new(HttpContractMapHandler)

			contractMapId := action.Trigger.After.ContractMapId
			httpContractMapHandler.ContractMap = settings.Contract.GetContractById(contractMapId)

			currentHandler.SetNext(httpContractMapHandler)
			currentHandler = httpContractMapHandler
		}

		currentHandler.SetNext(new(HttpLastHandler))

		chain.MapHandle[endpoint.Route] = firstHandler
	}
}

func (chain *Chain) GetByRoute(route string) Handler {
	return chain.MapHandle[route]
}
