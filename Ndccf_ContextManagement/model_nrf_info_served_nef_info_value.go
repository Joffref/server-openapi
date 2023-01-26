/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

type NrfInfoServedNefInfoValue struct {

	// Identity of the NEF
	NefId string `json:"nefId,omitempty"`

	PfdData PfdData `json:"pfdData,omitempty"`

	AfEeData AfEventExposureData `json:"afEeData,omitempty"`

	GpsiRanges []IdentityRange `json:"gpsiRanges,omitempty"`

	ExternalGroupIdentifiersRanges []IdentityRange `json:"externalGroupIdentifiersRanges,omitempty"`

	ServedFqdnList []string `json:"servedFqdnList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	DnaiList []string `json:"dnaiList,omitempty"`

	UnTrustAfInfoList []UnTrustAfInfo `json:"unTrustAfInfoList,omitempty"`

	UasNfFunctionalityInd bool `json:"uasNfFunctionalityInd,omitempty"`
}
