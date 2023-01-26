/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// AmfInfo - Information of an AMF NF Instance
type AmfInfo struct {

	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits). 
	AmfSetId string `json:"amfSetId"`

	// String identifying the AMF Set ID (10 bits) as specified in clause 2.10.1 of 3GPP TS 23.003.  It is encoded as a string of 3 hexadecimal characters where the first character is limited to  values 0 to 3 (i.e. 10 bits) 
	AmfRegionId string `json:"amfRegionId"`

	GuamiList []Guami `json:"guamiList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	BackupInfoAmfFailure []Guami `json:"backupInfoAmfFailure,omitempty"`

	BackupInfoAmfRemoval []Guami `json:"backupInfoAmfRemoval,omitempty"`

	N2InterfaceAmfInfo N2InterfaceAmfInfo `json:"n2InterfaceAmfInfo,omitempty"`

	AmfOnboardingCapability bool `json:"amfOnboardingCapability,omitempty"`
}
