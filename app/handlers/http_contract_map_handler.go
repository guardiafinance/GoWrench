package handlers

import (
	"context"
	contexts "wrench/app/contexts"
	"wrench/app/json_map"
	"wrench/app/manifest/contract_settings/maps"
)

type HttpContractMapHandler struct {
	Next        Handler
	ContractMap *maps.ContractMapSetting
}

func (handler *HttpContractMapHandler) Do(ctx context.Context, wrenchContext *contexts.WrenchContext, bodyContext *contexts.BodyContext) {

	if wrenchContext.HasError == false {
		currentBodyContext := bodyContext.BodyArray

		if handler.ContractMap.Properties != nil {
			currentBodyContext = json_map.RenameProperties(currentBodyContext, handler.ContractMap.Properties)
		}

		if handler.ContractMap.Remove != nil {
			currentBodyContext = json_map.RemoveProperties(currentBodyContext, handler.ContractMap.Remove)
		}

		bodyContext.BodyArray = currentBodyContext
	}

	if handler.Next != nil {
		handler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (handler *HttpContractMapHandler) SetNext(next Handler) {
	handler.Next = next
}
