/*
 * GMDviaMBMSbyMB2
 *
 * API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_GMDviaMBMSbyMB2

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
