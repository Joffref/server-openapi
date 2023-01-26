/*
 * Namf_Communication
 *
 * AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_Communication

import (
	"time"
)

// TimeWindow - Represents a time window identified by a start time and a stop time.
type TimeWindow struct {

	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime"`

	// string with format \"date-time\" as defined in OpenAPI.
	StopTime time.Time `json:"stopTime"`
}
