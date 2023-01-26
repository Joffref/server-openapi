/*
 * 3gpp-data-reporting
 *
 * API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReporting

import (
	"time"
)

type TripPlanRecordAllOf struct {

	StartingPoint LocationData `json:"startingPoint"`

	Waypoints []LocationData `json:"waypoints,omitempty"`

	Destination LocationData `json:"destination"`

	// Indicates value of horizontal speed.
	EstimatedAverageSpeed float32 `json:"estimatedAverageSpeed,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	EstimatedArrivalTime time.Time `json:"estimatedArrivalTime,omitempty"`
}
