/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// MbSmfInfo - Information of an MB-SMF NF Instance
type MbSmfInfo struct {

	// A map (list of key-value pairs) where a valid JSON string serves as key
	SNssaiInfoList map[string]SnssaiMbSmfInfoItem `json:"sNssaiInfoList,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	TmgiRangeList map[string]TmgiRange `json:"tmgiRangeList,omitempty"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	// A map (list of key-value pairs) where a valid JSON string serves as key
	MbsSessionList map[string]MbsSession `json:"mbsSessionList,omitempty"`
}
