/*
 * Nnwdaf_DataManagement
 *
 * Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nnwdaf_DataManagement

// MbsServiceArea - MBS Service Area
type MbsServiceArea struct {

	// List of NR cell Ids
	NcgiList []NcgiTai `json:"ncgiList,omitempty"`

	// List of tracking area Ids
	TaiList []Tai `json:"taiList,omitempty"`
}
