/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// SpatialValidityCond1 - Contains the Spatial Validity Condition.
type SpatialValidityCond1 struct {

	TrackingAreaList []Tai1 `json:"trackingAreaList,omitempty"`

	Countries []string `json:"countries,omitempty"`

	GeographicalServiceArea GeoServiceArea1 `json:"geographicalServiceArea,omitempty"`
}
