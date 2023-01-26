/*
 * Nmfaf_3daDataManagement
 *
 * MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3daDataManagement

// ReportingOptions - Represents reporting options for processed notifications.
type ReportingOptions struct {

	NotifyWindow TimeWindow `json:"notifyWindow,omitempty"`

	// indicating a time in seconds.
	NotifyPeriod int32 `json:"notifyPeriod,omitempty"`

	// indicating a time in seconds.
	NotifyPeriodInc int32 `json:"notifyPeriodInc,omitempty"`

	// Notifications for the present subscription are sent only upon occurrence of events of the subscription with identifier that matches this attribute. 
	DepEventSubId string `json:"depEventSubId,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MinClubbedNotif int32 `json:"minClubbedNotif,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxClubbedNotif int32 `json:"maxClubbedNotif,omitempty"`
}