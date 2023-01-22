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

// CommunicationCollection - Contains communication information.
type CommunicationCollection struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTime time.Time `json:"endTime"`

	// Unsigned integer identifying a volume in units of bytes.
	UlVol int64 `json:"ulVol"`

	// Unsigned integer identifying a volume in units of bytes.
	DlVol int64 `json:"dlVol"`
}