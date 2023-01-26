/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// UeAccessBehaviorReportItem - Report Item for UE Access Behavior Trends event.
type UeAccessBehaviorReportItem struct {

	StateTransitionType AccessStateTransitionType `json:"stateTransitionType"`

	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`

	// indicating a time in seconds.
	Duration int32 `json:"duration"`
}
