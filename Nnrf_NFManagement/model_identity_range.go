/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.3.0-alpha.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnrf_NFManagement

// IdentityRange - A range of GPSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching 
type IdentityRange struct {

	Start string `json:"start,omitempty"`

	End string `json:"end,omitempty"`

	Pattern string `json:"pattern,omitempty"`
}
