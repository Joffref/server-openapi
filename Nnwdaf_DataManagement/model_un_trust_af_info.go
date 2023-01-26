/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// UnTrustAfInfo - Information of a untrusted AF Instance
type UnTrustAfInfo struct {

	AfId string `json:"afId"`

	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`

	MappingInd bool `json:"mappingInd,omitempty"`
}
