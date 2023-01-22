/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package CommonData

import (
	"time"
)

// VolumeTimedReport - Contains Usage data information.
type VolumeTimedReport struct {

	// string with format 'date-time' as defined in OpenAPI.
	StartTimeStamp time.Time `json:"startTimeStamp"`

	// string with format 'date-time' as defined in OpenAPI.
	EndTimeStamp time.Time `json:"endTimeStamp"`

	// string with format 'int64' as defined in OpenAPI.
	DownlinkVolume int64 `json:"downlinkVolume"`

	// string with format 'int64' as defined in OpenAPI.
	UplinkVolume int64 `json:"uplinkVolume"`
}