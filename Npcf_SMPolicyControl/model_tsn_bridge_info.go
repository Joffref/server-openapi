/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

// TsnBridgeInfo - Contains parameters that describe and identify the TSC user plane node.
type TsnBridgeInfo struct {

	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	BridgeId int32 `json:"bridgeId,omitempty"`

	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	DsttAddr string `json:"dsttAddr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttPortNum int32 `json:"dsttPortNum,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	DsttResidTime int32 `json:"dsttResidTime,omitempty"`
}
