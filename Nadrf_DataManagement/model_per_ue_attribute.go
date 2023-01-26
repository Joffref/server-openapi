/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

import (
	"time"
)

// PerUeAttribute - UE application data collected per UE.
type PerUeAttribute struct {

	UeDest LocationArea5G `json:"ueDest,omitempty"`

	Route string `json:"route,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	AvgSpeed string `json:"avgSpeed,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeOfArrival time.Time `json:"timeOfArrival,omitempty"`
}