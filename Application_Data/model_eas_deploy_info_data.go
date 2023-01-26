/*
 * Unified Data Repository Service API file for Application Data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Application_Data

// EasDeployInfoData - Represents the EAS Deployment Information to be reported.
type EasDeployInfoData struct {

	AppId string `json:"appId,omitempty"`

	// list of DNS server identifier (consisting of IP address and port) and/or IP address(s)  of the EAS in the local DN for each DNAI. The key of map is the DNAI. 
	DnaiInfos map[string]DnaiInformation `json:"dnaiInfos,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	FqdnPatternList []FqdnPatternMatchingRule `json:"fqdnPatternList"`

	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	InternalGroupId string `json:"internalGroupId,omitempty"`

	Snssai Snssai `json:"snssai,omitempty"`
}
