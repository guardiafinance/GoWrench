package handlers

import (
	"context"
	"fmt"
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
		value, _ := json_map.GetValue(bodyContext.BodyArray, "name")
		value2, _ := json_map.GetValue(bodyContext.BodyArray, "address.street")

		fmt.Print(value)
		fmt.Print(value2)
	}

	if handler.Next != nil {
		handler.Next.Do(ctx, wrenchContext, bodyContext)
	}
}

func (handler *HttpContractMapHandler) SetNext(next Handler) {
	handler.Next = next
}
