/*
 * 3gpp-mbs-session
 *
 * API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MBSSession

// NcgiTai - List of NR cell ids, with their pertaining TAIs
type NcgiTai struct {

	Tai Tai `json:"tai"`

	// List of List of NR cell ids
	CellList []Ncgi `json:"cellList"`
}
