/*
 * Nmbsmf-MBSSession
 *
 * MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 
 *
 * API version: 1.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nmbsmf_MBSSession

// ExternalMbsServiceArea - List of geographic area or list of civic address info for MBS Service Area
type ExternalMbsServiceArea struct {

	GeographicAreaList []GeographicArea `json:"geographicAreaList,omitempty"`

	CivicAddressList []CivicAddress `json:"civicAddressList,omitempty"`
}
