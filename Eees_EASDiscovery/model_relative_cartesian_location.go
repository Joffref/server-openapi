/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EASDiscovery

// RelativeCartesianLocation - Relative Cartesian Location
type RelativeCartesianLocation struct {

	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x"`

	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y"`

	// string with format 'float' as defined in OpenAPI.
	Z float32 `json:"z,omitempty"`
}