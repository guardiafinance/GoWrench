package cross_cutting

import (
	"wrench/app/manifest/application_settings"
	"wrench/app/manifest/contract_settings/maps"
	"wrench/app/manifest/token_credential_settings"
)

func GetContractById(contractId string) *maps.ContractMapSetting {
	appSetting := application_settings.ApplicationSettingsStatic

	if appSetting.Contract == nil {
		return nil
	}

	if appSetting.Contract.Maps == nil ||
		len(appSetting.Contract.Maps) == 0 {
		return nil
	}

	var contractMapSetting *maps.ContractMapSetting = nil
	for _, contractMap := range appSetting.Contract.Maps {
		if contractMap.Id == contractId {
			contractMapSetting = contractMap
			break
		}
	}

	return contractMapSetting
}

func GetTokenCredentialById(tokenCredentialId string) *token_credential_settings.TokenCredentialSetting {
	appSetting := application_settings.ApplicationSettingsStatic

	if appSetting.TokenCredentials == nil {
		return nil
	}

	if len(appSetting.TokenCredentials) == 0 {
		return nil
	}

	var tokenCredential *token_credential_settings.TokenCredentialSetting = nil
	for _, tokenCredentialSetting := range appSetting.TokenCredentials {
		if tokenCredentialSetting.Id == tokenCredentialId {
			tokenCredential = tokenCredentialSetting
			break
		}
	}

	return tokenCredential
}
