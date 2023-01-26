/*
 * Ndccf_ContextManagement
 *
 * DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_ContextManagement

// NgApCause - Represents the NGAP cause.
type NgApCause struct {

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Group int32 `json:"group"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Value int32 `json:"value"`
}
