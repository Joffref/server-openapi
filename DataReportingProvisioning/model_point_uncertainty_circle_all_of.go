/*
 * 3gpp-data-reporting-provisioning
 *
 * API for 3GPP Data Reporting and Provisioning.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_DataReportingProvisioning

type PointUncertaintyCircleAllOf struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of uncertainty.
	Uncertainty float32 `json:"uncertainty"`
}
