/*
 * Nudm_PP
 *
 * Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_PP

// GeoServiceArea - List of geographic area or list of civic address info
type GeoServiceArea struct {

	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`

	CivicAddressList []CivicAddress `json:"civicAddressList,omitempty"`
}
