/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// SmfInfo - Information of an SMF NF Instance
type SmfInfo struct {

	SNssaiSmfInfoList []SnssaiSmfInfoItem `json:"sNssaiSmfInfoList"`

	TaiList []Tai `json:"taiList,omitempty"`

	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`

	// Fully Qualified Domain Name
	PgwFqdn string `json:"pgwFqdn,omitempty"`

	PgwIpAddrList []IpAddr `json:"pgwIpAddrList,omitempty"`

	AccessType []AccessType `json:"accessType,omitempty"`

	Priority int32 `json:"priority,omitempty"`

	VsmfSupportInd bool `json:"vsmfSupportInd,omitempty"`

	PgwFqdnList []string `json:"pgwFqdnList,omitempty"`

	// Deprecated
	SmfOnboardingCapability bool `json:"smfOnboardingCapability,omitempty"`

	IsmfSupportInd bool `json:"ismfSupportInd,omitempty"`

	SmfUPRPCapability bool `json:"smfUPRPCapability,omitempty"`
}
