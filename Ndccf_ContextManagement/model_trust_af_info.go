/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// TrustAfInfo - Information of a trusted AF Instance
type TrustAfInfo struct {

	SNssaiInfoList []SnssaiInfoItem `json:"sNssaiInfoList,omitempty"`

	AfEvents []AfEvent `json:"afEvents,omitempty"`

	AppIds []string `json:"appIds,omitempty"`

	InternalGroupId []string `json:"internalGroupId,omitempty"`

	MappingInd bool `json:"mappingInd,omitempty"`
}