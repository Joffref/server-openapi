/*
 * Ngmlc_Location
 *
 * GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ngmlc_Location

// AreaEventInfo - Indicates the information of area based event reporting.
type AreaEventInfo struct {

	AreaDefinition []ReportingArea `json:"areaDefinition"`

	OccurrenceInfo OccurrenceInfo `json:"occurrenceInfo,omitempty"`

	// Minimum interval between event reports.
	MinimumInterval int32 `json:"minimumInterval,omitempty"`

	// Maximum interval between event reports.
	MaximumInterval int32 `json:"maximumInterval,omitempty"`

	// Maximum time interval between consecutive evaluations by a UE of a trigger event.
	SamplingInterval int32 `json:"samplingInterval,omitempty"`

	// Maximum duration of event reporting.
	ReportingDuration int32 `json:"reportingDuration,omitempty"`

	ReportingLocationReq bool `json:"reportingLocationReq,omitempty"`
}
