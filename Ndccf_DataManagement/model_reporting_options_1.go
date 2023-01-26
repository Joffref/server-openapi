/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

import (
	"time"
)

type ReportingOptions1 struct {

	ReportMode EventReportMode `json:"reportMode,omitempty"`

	MaxNumOfReports int32 `json:"maxNumOfReports,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	Expiry time.Time `json:"expiry,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	SamplingRatio int32 `json:"samplingRatio,omitempty"`

	// indicating a time in seconds.
	GuardTime int32 `json:"guardTime,omitempty"`

	// indicating a time in seconds.
	ReportPeriod int32 `json:"reportPeriod,omitempty"`

	NotifFlag NotificationFlag `json:"notifFlag,omitempty"`
}
