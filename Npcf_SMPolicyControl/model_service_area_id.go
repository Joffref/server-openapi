/*
 * Npcf_SMPolicyControl API
 *
 * Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Npcf_SMPolicyControl

// ServiceAreaId - Contains a Service Area Identifier as defined in 3GPP TS 23.003, clause 12.5.
type ServiceAreaId struct {

	PlmnId PlmnId `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac"`

	// Service Area Code.
	Sac string `json:"sac"`
}
