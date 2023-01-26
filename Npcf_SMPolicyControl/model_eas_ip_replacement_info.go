/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

// EasIpReplacementInfo - Contains EAS IP replacement information for a Source and a Target EAS.
type EasIpReplacementInfo struct {

	Source EasServerAddress `json:"source"`

	Target EasServerAddress `json:"target"`
}
