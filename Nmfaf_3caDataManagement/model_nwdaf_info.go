/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// NwdafInfo - Information of a NWDAF NF Instance
type NwdafInfo struct {

	EventIds []EventId `json:"eventIds,omitempty"`

	NwdafEvents []NwdafEvent `json:"nwdafEvents,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	NwdafCapability NwdafCapability `json:"nwdafCapability,omitempty"`

	// indicating a time in seconds.
	AnalyticsDelay int32 `json:"analyticsDelay,omitempty"`

	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`

	ServingNfTypeList []NfType `json:"servingNfTypeList,omitempty"`

	MlAnalyticsList []MlAnalyticsInfo `json:"mlAnalyticsList,omitempty"`
}
