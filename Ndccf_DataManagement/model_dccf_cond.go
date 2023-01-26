/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// DccfCond - Subscription to a set of NF Instances (DCCFs), identified by NF types, NF Set Id(s) or DCCF Serving Area information, i.e. list of TAIs served by the DCCF 
type DccfCond struct {

	ConditionType string `json:"conditionType"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	ServingNfTypeList []NfType `json:"servingNfTypeList,omitempty"`

	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`
}
