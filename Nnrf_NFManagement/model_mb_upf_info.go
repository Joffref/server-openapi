/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

// MbUpfInfo - Information of an MB-UPF NF Instance
type MbUpfInfo struct {

	SNssaiMbUpfInfoList []SnssaiUpfInfoItem `json:"sNssaiMbUpfInfoList"`

	MbSmfServingArea []string `json:"mbSmfServingArea,omitempty"`

	InterfaceMbUpfInfoList []InterfaceUpfInfoItem `json:"interfaceMbUpfInfoList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	SupportedPfcpFeatures string `json:"supportedPfcpFeatures,omitempty"`
}
