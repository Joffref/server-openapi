/*
 * Nudr_DataRepository API OpenAPI file
 *
 * Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi_Nudr_DR

// RouteInformation - At least one of the \"ipv4Addr\" attribute and the \"ipv6Addr\" attribute shall be included in the \"RouteInformation\" data type.  
type RouteInformation struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}
