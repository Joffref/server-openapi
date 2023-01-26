/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// RelativeCartesianLocation - Relative Cartesian Location
type RelativeCartesianLocation struct {

	// string with format 'float' as defined in OpenAPI.
	X float32 `json:"x"`

	// string with format 'float' as defined in OpenAPI.
	Y float32 `json:"y"`

	// string with format 'float' as defined in OpenAPI.
	Z float32 `json:"z,omitempty"`
}