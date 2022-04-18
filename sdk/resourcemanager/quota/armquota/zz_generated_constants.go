//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

const (
	moduleName    = "armquota"
	moduleVersion = "v0.4.0"
)

// LimitType - The limit object type.
type LimitType string

const (
	LimitTypeLimitValue LimitType = "LimitValue"
)

// PossibleLimitTypeValues returns the possible values for the LimitType const type.
func PossibleLimitTypeValues() []LimitType {
	return []LimitType{
		LimitTypeLimitValue,
	}
}

// QuotaLimitTypes - The quota or usages limit types.
type QuotaLimitTypes string

const (
	QuotaLimitTypesIndependent QuotaLimitTypes = "Independent"
	QuotaLimitTypesShared      QuotaLimitTypes = "Shared"
)

// PossibleQuotaLimitTypesValues returns the possible values for the QuotaLimitTypes const type.
func PossibleQuotaLimitTypesValues() []QuotaLimitTypes {
	return []QuotaLimitTypes{
		QuotaLimitTypesIndependent,
		QuotaLimitTypesShared,
	}
}

// QuotaRequestState - Quota request status.
type QuotaRequestState string

const (
	QuotaRequestStateAccepted   QuotaRequestState = "Accepted"
	QuotaRequestStateFailed     QuotaRequestState = "Failed"
	QuotaRequestStateInProgress QuotaRequestState = "InProgress"
	QuotaRequestStateInvalid    QuotaRequestState = "Invalid"
	QuotaRequestStateSucceeded  QuotaRequestState = "Succeeded"
)

// PossibleQuotaRequestStateValues returns the possible values for the QuotaRequestState const type.
func PossibleQuotaRequestStateValues() []QuotaRequestState {
	return []QuotaRequestState{
		QuotaRequestStateAccepted,
		QuotaRequestStateFailed,
		QuotaRequestStateInProgress,
		QuotaRequestStateInvalid,
		QuotaRequestStateSucceeded,
	}
}

// UsagesTypes - The quota or usages limit types.
type UsagesTypes string

const (
	UsagesTypesCombined   UsagesTypes = "Combined"
	UsagesTypesIndividual UsagesTypes = "Individual"
)

// PossibleUsagesTypesValues returns the possible values for the UsagesTypes const type.
func PossibleUsagesTypesValues() []UsagesTypes {
	return []UsagesTypes{
		UsagesTypesCombined,
		UsagesTypesIndividual,
	}
}
