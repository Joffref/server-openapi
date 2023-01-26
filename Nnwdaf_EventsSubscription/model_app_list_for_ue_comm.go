/*
 * Nnwdaf_EventsSubscription
 *
 * Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_EventsSubscription

import (
	"time"
)

// AppListForUeComm - Represents the analytics of the application list used by UE.
type AppListForUeComm struct {

	// String providing an application identifier.
	AppId string `json:"appId"`

	// string with format 'date-time' as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// indicating a time in seconds.
	AppDur int32 `json:"appDur,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	OccurRatio int32 `json:"occurRatio,omitempty"`

	SpatialValidity NetworkAreaInfo `json:"spatialValidity,omitempty"`
}
