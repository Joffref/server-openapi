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

type ReachabilityReport struct {

	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AmfInstanceId string `json:"amfInstanceId,omitempty"`

	AccessTypeList []AccessType `json:"accessTypeList,omitempty"`

	Reachability UeReachability `json:"reachability,omitempty"`

	// string with format 'date-time' as defined in OpenAPI.
	MaxAvailabilityTime time.Time `json:"maxAvailabilityTime,omitempty"`

	IdleStatusIndication IdleStatusIndication `json:"idleStatusIndication,omitempty"`
}
