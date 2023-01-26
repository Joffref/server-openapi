/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type NrfInfoServedPcfInfoValue struct {

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	DnnList []string `json:"dnnList,omitempty"`

	SupiRanges []SupiRange `json:"supiRanges,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	// Fully Qualified Domain Name
	RxDiamHost string `json:"rxDiamHost,omitempty"`

	// Fully Qualified Domain Name
	RxDiamRealm string `json:"rxDiamRealm,omitempty"`

	V2xSupportInd bool `json:"v2xSupportInd,omitempty"`

	ProseSupportInd bool `json:"proseSupportInd,omitempty"`

	ProseCapability ProSeCapability `json:"proseCapability,omitempty"`

	V2xCapability V2xCapability `json:"v2xCapability,omitempty"`
}
