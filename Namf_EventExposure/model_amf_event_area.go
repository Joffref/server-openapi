/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_EventExposure

// AmfEventArea - Represents an area to be monitored by an AMF event
type AmfEventArea struct {

	PresenceInfo PresenceInfo `json:"presenceInfo,omitempty"`

	LadnInfo LadnInfo `json:"ladnInfo,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	// Contains the Identifier of the selected Network Slice instance
	NsiId string `json:"nsiId,omitempty"`
}
