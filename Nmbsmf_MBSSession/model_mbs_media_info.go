/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// MbsMediaInfo - Represent MBS Media Information.
type MbsMediaInfo struct {

	MbsMedType MediaType `json:"mbsMedType,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxReqMbsBwDl string `json:"maxReqMbsBwDl,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MinReqMbsBwDl string `json:"minReqMbsBwDl,omitempty"`

	Codecs []string `json:"codecs,omitempty"`
}
