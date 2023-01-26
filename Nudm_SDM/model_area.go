/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudm_SDM

// Area - Provides area information.
type Area struct {

	Tacs []string `json:"tacs,omitempty"`

	// Values are operator specific.
	AreaCode string `json:"areaCode,omitempty"`
}
