/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_PfdManagement

// EllipsoidArc - Ellipsoid Arc.
type EllipsoidArc struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of the inner radius.
	InnerRadius int32 `json:"innerRadius"`

	// Indicates value of uncertainty.
	UncertaintyRadius float32 `json:"uncertaintyRadius"`

	// Indicates value of angle.
	OffsetAngle int32 `json:"offsetAngle"`

	// Indicates value of angle.
	IncludedAngle int32 `json:"includedAngle"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
