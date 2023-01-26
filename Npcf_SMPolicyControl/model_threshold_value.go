/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

// ThresholdValue - Indicates the threshold value(s) for RTT and/or Packet Loss Rate.
type ThresholdValue struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	RttThres *int32 `json:"rttThres,omitempty"`

	// This data type is defined in the same way as the 'PacketLossRate' data type, but with the OpenAPI 'nullable: true' property. 
	PlrThres *int32 `json:"plrThres,omitempty"`
}
