/*
 * Common Data Types
 *
 * Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   
 *
 * API version: 1.4.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_CommonData

// PointUncertaintyEllipse - Ellipsoid point with uncertainty ellipse.
type PointUncertaintyEllipse struct {

	Point GeographicalCoordinates `json:"point"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
