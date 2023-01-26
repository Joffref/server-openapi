/*
 * Unified Data Repository Service API file for structured data for exposure
 *
 * The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: -
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Exposure_Data

// RouteInformation - At least one of the \"ipv4Addr\" attribute and the \"ipv6Addr\" attribute shall be included in the \"RouteInformation\" data type.  
type RouteInformation struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}
