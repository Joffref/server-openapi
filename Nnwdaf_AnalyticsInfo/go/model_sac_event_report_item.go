/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nnwdaf_AnalyticsInfo

import (
	"time"
)

// SacEventReportItem - Represents a report triggered by a subscribed event type
type SacEventReportItem struct {

	EventType SacEventType `json:"eventType"`

	EventState SacEventState `json:"eventState"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`

	EventFilter Snssai `json:"eventFilter"`

	SliceStautsInfo SacEventStatus `json:"sliceStautsInfo,omitempty"`
}