/*
 * 3gpp-am-policyauthorization
 *
 * API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AMPolicyAuthorization

type PointAltitudeAllOf struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}
