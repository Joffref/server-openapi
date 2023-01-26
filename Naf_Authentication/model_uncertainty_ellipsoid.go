/*
 * Naf_Authentication
 *
 * AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Naf_Authentication

// UncertaintyEllipsoid - Ellipsoid with uncertainty
type UncertaintyEllipsoid struct {

	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`

	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`

	// Indicates value of uncertainty.
	Vertical float32 `json:"vertical"`

	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}
