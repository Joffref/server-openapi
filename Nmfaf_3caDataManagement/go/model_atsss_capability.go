/*
 * Nmfaf_3caDataManagement
 *
 * MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package Nmfaf_3caDataManagement

// AtsssCapability - Containes Capability to support procedures related to Access Traffic Steering, Switching, Splitting. 
type AtsssCapability struct {

	// Indicates the ATSSS-LL capability to support procedures related to Access Traffic Steering, Switching, Splitting (see clauses 4.2.10, 5.32 of 3GPP TS 23.501). true: Supported false (default): Not Supported 
	AtsssLL bool `json:"atsssLL,omitempty"`

	// Indicates the MPTCP capability to support procedures related to Access Traffic Steering, Switching, Splitting (see clauses 4.2.10, 5.32 of 3GPP TS 23.501 true: Supported false (default): Not Supported 
	Mptcp bool `json:"mptcp,omitempty"`

	// This IE is only used by the UPF to indicate whether the UPF supports RTT measurement without PMF (see clauses 5.32.2, 6.3.3.3 of 3GPP TS 23.501 true: Supported false (default): Not Supported 
	RttWithoutPmf bool `json:"rttWithoutPmf,omitempty"`
}