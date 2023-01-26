/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

import (
	"time"
)

// AfEventNotification - Represents information related to an event to be reported.
type AfEventNotification struct {

	Event AfEvent `json:"event"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	SvcExprcInfos []ServiceExperienceInfoPerApp `json:"svcExprcInfos,omitempty"`

	UeMobilityInfos []UeMobilityCollection `json:"ueMobilityInfos,omitempty"`

	UeCommInfos []UeCommunicationCollection `json:"ueCommInfos,omitempty"`

	ExcepInfos []ExceptionInfo `json:"excepInfos,omitempty"`

	CongestionInfos []UserDataCongestionCollection `json:"congestionInfos,omitempty"`

	PerfDataInfos []PerformanceDataCollection `json:"perfDataInfos,omitempty"`

	DispersionInfos []DispersionCollection1 `json:"dispersionInfos,omitempty"`

	CollBhvrInfs []CollectiveBehaviourInfo `json:"collBhvrInfs,omitempty"`

	MsQoeMetrInfos []MsQoeMetricsCollection `json:"msQoeMetrInfos,omitempty"`

	MsConsumpInfos []MsConsumptionCollection `json:"msConsumpInfos,omitempty"`

	MsNetAssInvInfos []MsNetAssInvocationCollection `json:"msNetAssInvInfos,omitempty"`

	MsDynPlyInvInfos []MsDynPolicyInvocationCollection `json:"msDynPlyInvInfos,omitempty"`

	MsAccActInfos []MsAccessActivityCollection `json:"msAccActInfos,omitempty"`
}
