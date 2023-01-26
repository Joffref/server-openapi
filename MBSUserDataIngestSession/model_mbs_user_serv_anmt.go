/*
 * 3gpp-mbs-ud-ingest
 *
 * API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSUserDataIngestSession

import (
	"time"
)

// MbsUserServAnmt - Represents the MBS User Service Announcement currently associated with the MBS User Data  Ingest Session. 
type MbsUserServAnmt struct {

	ExtServiceId []string `json:"extServiceId"`

	ServClass string `json:"servClass"`

	// string with format \"date-time\" as defined in OpenAPI.
	StartTime time.Time `json:"startTime,omitempty"`

	// string with format \"date-time\" as defined in OpenAPI.
	EndTime time.Time `json:"endTime,omitempty"`

	ServNameDescs []ServiceNameDescription `json:"servNameDescs"`

	MainServLang string `json:"mainServLang,omitempty"`

	// Represents the set of MBS Distribution Session Announcements currently associated with  this MBS User Service Announcement. 
	MbsDistSessAnmt map[string]MbsDistSessionAnmt `json:"mbsDistSessAnmt,omitempty"`
}
