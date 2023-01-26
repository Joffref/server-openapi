/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_EventExposure

import (
	"time"
)

// IdleStatusIndication - Represents the idle status indication.
type IdleStatusIndication struct {

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp,omitempty"`

	// indicating a time in seconds.
	ActiveTime int32 `json:"activeTime,omitempty"`

	// indicating a time in seconds.
	SubsRegTimer int32 `json:"subsRegTimer,omitempty"`

	EdrxCycleLength int32 `json:"edrxCycleLength,omitempty"`

	SuggestedNumOfDlPackets int32 `json:"suggestedNumOfDlPackets,omitempty"`
}
