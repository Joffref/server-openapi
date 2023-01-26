/*
 * Nadrf_DataManagement
 *
 * ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nadrf_DataManagement

// FormattingInstruction - Contains data or analytics formatting instructions.
type FormattingInstruction struct {

	// Indicates that notifications shall be buffered until the NF service consumer requests their delivery. 
	ConsTrigNotif bool `json:"consTrigNotif,omitempty"`

	ReportingOptions ReportingOptions1 `json:"reportingOptions,omitempty"`
}
