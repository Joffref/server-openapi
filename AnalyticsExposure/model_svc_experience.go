/*
 * 3gpp-analyticsexposure
 *
 * API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_AnalyticsExposure

// SvcExperience - Contains a mean opinion score with the customized range.
type SvcExperience struct {

	// string with format 'float' as defined in OpenAPI.
	Mos float32 `json:"mos,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	UpperRange float32 `json:"upperRange,omitempty"`

	// string with format 'float' as defined in OpenAPI.
	LowerRange float32 `json:"lowerRange,omitempty"`
}
