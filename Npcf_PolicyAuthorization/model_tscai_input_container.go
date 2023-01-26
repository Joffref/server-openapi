/*
 * Npcf_PolicyAuthorization Service API
 *
 * PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_PolicyAuthorization

import (
	"time"
)

// TscaiInputContainer - Indicates TSC Traffic pattern.
type TscaiInputContainer struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Periodicity int32 `json:"periodicity,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	BurstArrivalTime time.Time `json:"burstArrivalTime,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInNumMsg int32 `json:"surTimeInNumMsg,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	SurTimeInTime int32 `json:"surTimeInTime,omitempty"`
}