/*
 * VAE_HDMapDynamicInfo
 *
 * API for VAE HDMapDynamicInfo Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_VAE_HDMapDynamicInfo

// LocationAreaId - Contains a Location area identification as defined in 3GPP TS 23.003, clause 4.1.
type LocationAreaId struct {

	PlmnId PlmnId `json:"plmnId"`

	// Location Area Code.
	Lac string `json:"lac"`
}
