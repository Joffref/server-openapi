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

// UeTrajectoryCollection - Contains UE trajectory information associated with an application.
type UeTrajectoryCollection struct {

	// string with format 'date-time' as defined in OpenAPI.
	Ts time.Time `json:"ts"`

	LocArea LocationArea5G `json:"locArea"`
}
