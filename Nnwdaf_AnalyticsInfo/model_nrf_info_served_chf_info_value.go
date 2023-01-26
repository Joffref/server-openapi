/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

type NrfInfoServedChfInfoValue struct {

	SupiRangeList []SupiRange `json:"supiRangeList,omitempty"`

	GpsiRangeList []IdentityRange `json:"gpsiRangeList,omitempty"`

	PlmnRangeList []PlmnRange `json:"plmnRangeList,omitempty"`

	// Identifier of a group of NFs.
	GroupId string `json:"groupId,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PrimaryChfInstance string `json:"primaryChfInstance,omitempty"`

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	SecondaryChfInstance string `json:"secondaryChfInstance,omitempty"`
}
