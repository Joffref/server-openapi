/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

import (
	"time"
)

// EventReportingRequirement - Represents the type of reporting that the subscription requires.
type EventReportingRequirement struct {

	Accuracy Accuracy `json:"accuracy,omitempty"`

	// Each element indicates the preferred accuracy level per analytics subset. It may be present if the \"listOfAnaSubsets\" attribute is present in the subscription request when the subscription event is NF_LOAD, UE_COMM, DISPERSION, NETWORK_PERFORMANCE, WLAN_PERFORMANCE, DN_PERFORMANCE or SERVICE_EXPERIENCE. 
	AccPerSubset []Accuracy `json:"accPerSubset,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTs time.Time `json:"startTs,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTs time.Time `json:"endTs,omitempty"`

	// Offset period in units of seconds to the reporting time, if the value is negative means  statistics in the past offset period, otherwise a positive value means prediction in the  future offset period. May be present if the \"repPeriod\" attribute is included within the  \"evtReq\" attribute. 
	OffsetPeriod int32 `json:"offsetPeriod,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SampRatio int32 `json:"sampRatio,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxObjectNbr int32 `json:"maxObjectNbr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxSupiNbr int32 `json:"maxSupiNbr,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeAnaNeeded time.Time `json:"timeAnaNeeded,omitempty"`

	AnaMeta []AnalyticsMetadata `json:"anaMeta,omitempty"`

	AnaMetaInd AnalyticsMetadataIndication `json:"anaMetaInd,omitempty"`
}
