/*
 * Ndccf_DataManagement
 *
 * DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Ndccf_DataManagement

// SdRange - A range of SDs (Slice Differentiators)
type SdRange struct {

	// First value identifying the start of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2. 
	Start string `json:"start,omitempty"`

	// Last value identifying the end of an SD range. This string shall be formatted as specified for the sd attribute of the Snssai data type in clause 5.4.4.2. 
	End string `json:"end,omitempty"`
}
