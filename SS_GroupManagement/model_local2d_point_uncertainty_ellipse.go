/*
 * SS_GroupManagement
 *
 * API for SEAL Group management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_SS_GroupManagement

// Local2dPointUncertaintyEllipse - Local 2D point with uncertainty ellipse
type Local2dPointUncertaintyEllipse struct {

	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
