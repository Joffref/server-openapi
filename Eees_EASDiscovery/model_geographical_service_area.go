/*
 * Eees_EASDiscovery
 *
 * API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Eees_EASDiscovery

// GeographicalServiceArea - Represents geographical service area information.
type GeographicalServiceArea struct {

	// A list of geographic area information.
	GeoArs []GeographicArea `json:"geoArs,omitempty"`

	// A list of civic address information.
	CivicAddrs []CivicAddress `json:"civicAddrs,omitempty"`
}
