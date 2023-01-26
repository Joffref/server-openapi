/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmfaf_3caDataManagement

// TopApplication - Top application that contributes the most to the traffic.
type TopApplication struct {

	// String providing an application identifier.
	AppId string `json:"appId,omitempty"`

	IpTrafficFilter FlowInfo `json:"ipTrafficFilter,omitempty"`

	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.  
	Ratio int32 `json:"ratio,omitempty"`
}
