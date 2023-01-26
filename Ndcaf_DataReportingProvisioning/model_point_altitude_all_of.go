/*
 * Ndcaf_DataReportingProvisioning
 *
 * Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndcaf_DataReportingProvisioning

type PointAltitudeAllOf struct {

	Point GeographicalCoordinates `json:"point"`

	// Indicates value of altitude.
	Altitude float64 `json:"altitude"`
}
