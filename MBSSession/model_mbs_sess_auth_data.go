/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSSession

// MbsSessAuthData - Represents the MBS Session Authorization data.
type MbsSessAuthData struct {

	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExtGroupId string `json:"extGroupId"`

	// Represents the list of the GPSI(s) of the member UE(s) constituting the multicast MBS group. Any value of type can be used as a key of the map. 
	GpsisList map[string]string `json:"gpsisList,omitempty"`

	MbsSessionIdList *Model5MbsAuthorizationInfo `json:"mbsSessionIdList"`
}