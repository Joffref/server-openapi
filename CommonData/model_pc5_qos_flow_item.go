/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// Pc5QosFlowItem - Contains a PC5 QOS flow.
type Pc5QosFlowItem struct {

	// Unsigned integer representing a 5G QoS Identifier (see clause 5.7.2.1 of 3GPP TS 23.501, within the range 0 to 255. 
	Pqi int32 `json:"pqi"`

	Pc5FlowBitRates Pc5FlowBitRates `json:"pc5FlowBitRates,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Range int32 `json:"range,omitempty"`
}
