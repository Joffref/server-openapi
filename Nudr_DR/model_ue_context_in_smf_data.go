/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

type UeContextInSmfData struct {

	// A map (list of key-value pairs where PduSessionId serves as key) of PduSessions
	PduSessions map[string]PduSession `json:"pduSessions,omitempty"`

	PgwInfo []PgwInfo `json:"pgwInfo,omitempty"`

	EmergencyInfo EmergencyInfo `json:"emergencyInfo,omitempty"`
}
