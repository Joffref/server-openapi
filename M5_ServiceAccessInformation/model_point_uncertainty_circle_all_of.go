/*
 * M5_ServiceAccessInformation
 *
 * 5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 2.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_M5_ServiceAccessInformation

type PointUncertaintyCircleAllOf struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}
