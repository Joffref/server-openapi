/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 1.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nsmf_EventExposure

// PlmnId - When PlmnId needs to be converted to string (e.g. when used in maps as key), the string  shall be composed of three digits \"mcc\" followed by \"-\" and two or three digits \"mnc\". 
type PlmnId struct {

	// Mobile Country Code part of the PLMN, comprising 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.  
	Mcc string `json:"mcc"`

	// Mobile Network Code part of the PLMN, comprising 2 or 3 digits, as defined in clause 9.3.3.5 of 3GPP TS 38.413.
	Mnc string `json:"mnc"`
}
