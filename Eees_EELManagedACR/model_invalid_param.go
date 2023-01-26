/*
 * EES EEL Managed ACR Service
 *
 * EES EEL Managed ACR Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EELManagedACR

// InvalidParam - Represents the description of invalid parameters, for a request rejected due to invalid parameters.
type InvalidParam struct {

	// Attribute's name encoded as a JSON Pointer, or header's name.
	Param string `json:"param"`

	// A human-readable reason, e.g. \"must be a positive integer\".
	Reason string `json:"reason,omitempty"`
}
