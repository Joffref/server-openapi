/*
 * Nnwdaf_AnalyticsInfo
 *
 * Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_AnalyticsInfo

import (
	"time"
)

// SmNasFromUe - Represents information on the SM NAS messages that SMF receives from UE for PDU Session. 
type SmNasFromUe struct {

	SmNasType string `json:"smNasType"`

	// string with format 'date-time' as defined in OpenAPI.
	TimeStamp time.Time `json:"timeStamp"`
}
