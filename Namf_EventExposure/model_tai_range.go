/*
 * Namf_EventExposure
 *
 * AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Namf_EventExposure

// TaiRange - Range of TAIs (Tracking Area Identities)
type TaiRange struct {

	PlmnId PlmnId `json:"plmnId"`

	TacRangeList []TacRange `json:"tacRangeList"`

	// This represents the Network Identifier, which together with a PLMN ID is used to identify an SNPN (see 3GPP TS 23.003 and 3GPP TS 23.501 clause 5.30.2.1).  
	Nid string `json:"nid,omitempty"`
}
