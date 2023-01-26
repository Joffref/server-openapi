/*
 * Unified Data Repository Service API file for policy data
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Policy_Data

// AmPolicyData - Contains the AM policy data for a given subscriber.
type AmPolicyData struct {

	// Contains Presence reporting area information. The praId attribute within the PresenceInfo data type is the key of the map. 
	PraInfos map[string]PresenceInfo `json:"praInfos,omitempty"`

	SubscCats []string `json:"subscCats,omitempty"`
}
