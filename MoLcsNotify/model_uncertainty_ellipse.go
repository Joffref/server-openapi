/*
 * 3gpp-mo-lcs-notify
 *
 * API for UE updated location information notification.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_MoLcsNotify

// UncertaintyEllipse - Ellipse with uncertainty.
type UncertaintyEllipse struct {

	// Indicates value of uncertainty.
	SemiMajor float32 `json:"semiMajor"`

	// Indicates value of uncertainty.
	SemiMinor float32 `json:"semiMinor"`

	// Indicates value of orientation angle.
	OrientationMajor int32 `json:"orientationMajor"`
}
