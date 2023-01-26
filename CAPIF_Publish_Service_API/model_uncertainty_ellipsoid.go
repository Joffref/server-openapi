/*
 * CAPIF_Publish_Service_API
 *
 * API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CAPIF_Publish_Service_API

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
