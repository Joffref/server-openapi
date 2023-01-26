/*
 * SS_LocationAreaInfoRetrieval
 *
 * API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_LocationAreaInfoRetrieval

import (
	"time"
)

// LmInformation - Represents the location information for a VAL User ID or a VAL UE ID.
type LmInformation struct {

	ValTgtUe ValTargetUe `json:"valTgtUe"`

	LocInfo LocationInfo `json:"locInfo"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp,omitempty"`

	// Identity of the VAL service
	ValSvcId string `json:"valSvcId,omitempty"`
}
