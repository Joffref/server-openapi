/*
 * 3gpp-pfd-management
 *
 * API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_PfdManagement

type Local2dPointUncertaintyEllipseAllOf struct {

	LocalOrigin LocalOrigin `json:"localOrigin"`

	Point RelativeCartesianLocation `json:"point"`

	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`

	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}
