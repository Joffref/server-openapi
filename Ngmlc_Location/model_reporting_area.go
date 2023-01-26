/*
 * Ngmlc_Location
 *
 * GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ngmlc_Location

// ReportingArea - Indicates an area for event reporting.
type ReportingArea struct {

	AreaType ReportingAreaType `json:"areaType"`

	Tai Tai `json:"tai,omitempty"`

	Ecgi Ecgi `json:"ecgi,omitempty"`

	Ncgi Ncgi `json:"ncgi,omitempty"`
}
