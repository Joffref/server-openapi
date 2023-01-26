/*
 * Npcf_MBSPolicyControl API
 *
 * MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_MBSPolicyControl

// MbsServiceInfo - Represent MBS Service Information.
type MbsServiceInfo struct {

	MbsMediaComps map[string]MbsMediaCompRm `json:"mbsMediaComps"`

	MbsSdfResPrio ReservPriority `json:"mbsSdfResPrio,omitempty"`

	// Contains an AF application identifier.
	AfAppId string `json:"afAppId,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MbsSessionAmbr string `json:"mbsSessionAmbr,omitempty"`
}
