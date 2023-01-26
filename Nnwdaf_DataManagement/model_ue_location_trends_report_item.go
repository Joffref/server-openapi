/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

import (
	"time"
)

// UeLocationTrendsReportItem - Report Item for UE Location Trends event.
type UeLocationTrendsReportItem struct {

	Tai Tai `json:"tai,omitempty"`

	Ncgi Ncgi `json:"ncgi,omitempty"`

	Ecgi Ecgi `json:"ecgi,omitempty"`

	N3gaLocation N3gaLocation `json:"n3gaLocation,omitempty"`

	// indicating a time in seconds.
	Spacing int32 `json:"spacing"`

	// indicating a time in seconds.
	Duration int32 `json:"duration"`

	// string with format 'date-time' as defined in OpenAPI.
	Timestamp time.Time `json:"timestamp"`
}
