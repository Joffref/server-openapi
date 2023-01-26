/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

import (
	"time"
)

// WlanPerTsPerformanceInfo - WLAN performance information per Time Slot during the analytics target period.
type WlanPerTsPerformanceInfo struct {

	// string with format 'date-time' as defined in OpenAPI.
	TsStart time.Time `json:"tsStart"`

	// indicating a time in seconds.
	TsDuration int32 `json:"tsDuration"`

	Rssi int32 `json:"rssi,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Rtt int32 `json:"rtt,omitempty"`

	TrafficInfo TrafficInformation `json:"trafficInfo,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	NumberOfUes int32 `json:"numberOfUes,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence int32 `json:"confidence,omitempty"`
}
